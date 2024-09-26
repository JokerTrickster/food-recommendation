package repository

import (
	"context"
	"fmt"
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"
	_redis "main/utils/db/redis"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRankingFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IRankingFoodRepository {
	return &RankingFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *RankingFoodRepository) FindAllRanking(ctx context.Context, redisKey string) ([]string, error) {
	// 내림차순으로 10개 멤버와 점수 가져오기
	foodList, err := g.RedisClient.ZRevRangeWithScores(ctx, redisKey, 0, 9).Result()
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 음식 이름과 점수를 저장할 맵 생성
	rankings := make([]string, 0)

	// 결과를 맵에 추가
	for _, z := range foodList {
		rankings = append(rankings, z.Member.(string))
	}

	return rankings, nil
}
func (g *RankingFoodRepository) FindPreviousRanking(ctx context.Context, todayRedisKey, yesterDayRedisKey string, food string, currentRank int) (string, error) {

	// 어제의 점수를 가져오기
	_, err := g.RedisClient.ZScore(ctx, yesterDayRedisKey, food).Result()
	if err == redis.Nil {
		return "new", nil // 이전 랭킹이 없으면 "new" 반환
	} else if err != nil {
		fmt.Println("Error fetching previous ranking:", err)
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 어제의 랭킹 가져오기
	prevRank, err := g.RedisClient.ZRevRank(ctx, yesterDayRedisKey, food).Result()
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 오늘의 랭킹 가져오기
	currentRankToday, err := g.RedisClient.ZRevRank(ctx, todayRedisKey, food).Result()
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 랭킹 변동 계산
	rankChange := int(prevRank) - int(currentRankToday)
	rankChangeStr := strconv.Itoa(rankChange)

	return rankChangeStr, nil
}

func (g *RankingFoodRepository) RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	//get rankings foods
	currentRankings, err := _redis.Client.ZRevRangeWithScores(ctx, _redis.RankingKey, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}
	result := []*entity.RankFoodRedis{}
	for _, z := range currentRankings {
		rankFood := &entity.RankFoodRedis{
			FoodID: z.Member.(string),
			Score:  z.Score,
		}
		result = append(result, rankFood)
	}

	return result, nil
}

func (g *RankingFoodRepository) PreviousRanking(ctx context.Context, food string) (int, error) {
	// Get previous ranking of food
	previousRank, err := _redis.Client.ZRevRank(ctx, _redis.PrevRankingKey, food).Result()
	if err != nil {
		if err == redis.Nil {
			return _redis.NewRank, nil
		}
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	return int(previousRank), nil
}

func (g *RankingFoodRepository) FindRankingTop10FoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	// gorm에서 food_histories 테이블에서 top 10가져오기
	var results []struct {
		FoodID int
		Count  int64
	}
	// SQL 쿼리 실행
	err := g.GormDB.Table("food_histories").
		Select("food_id, COUNT(food_id) as count").
		Group("food_id").
		Order("count DESC").
		Limit(10).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch top 10 food histories: %w", err)
	}

	// 결과에서 음식 이름 추출
	topFoods := make([]*entity.RankFoodRedis, 0)
	for _, r := range results {
		topFoods = append(topFoods, &entity.RankFoodRedis{
			FoodID: strconv.Itoa(r.FoodID),
			Score:  float64(r.Count),
		})
	}

	return topFoods, nil
}

func (g *RankingFoodRepository) IncrementFoodRanking(ctx context.Context, redisKey string, foodName string, score float64) error {
	// Increment food ranking in Redis
	_, err := _redis.Client.ZAdd(ctx, redisKey, redis.Z{Score: score, Member: foodName}).Result()
	if err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	return nil
}

func (g *RankingFoodRepository) PreviousRankingExist(ctx context.Context) (int, error) {
	// Check if previous ranking exists
	previousExist, err := _redis.Client.Exists(ctx, _redis.PrevRankingKey).Result()
	if err != nil {
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	return int(previousExist), nil
}

func (g *RankingFoodRepository) FindOneFoods(ctx context.Context, foodID int) (string, error) {
	// Find food name by food ID
	var foodName string
	err := g.GormDB.WithContext(ctx).Model(&mysql.Foods{}).Select("name").Where("id = ?", foodID).First(&foodName).Error
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromMysqlDB)
	}

	return foodName, nil
}

func (g *RankingFoodRepository) ExpireRanking(ctx context.Context, key string) error {
	// Set expiration time for key
	err := g.RedisClient.Expire(ctx, key, 30*time.Minute).Err() // RedisClient는 Redis 연결을 나타내는 변수입니다.
	if err != nil {
		return fmt.Errorf("failed to set expiration for key %s: %w", key, err)
	}

	return nil
}
