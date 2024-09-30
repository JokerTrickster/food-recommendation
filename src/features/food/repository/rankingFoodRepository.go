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

func (g *RankingFoodRepository) RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	//get rankings foods
	currentRankings, err := _redis.Client.ZRevRangeWithScores(ctx, _redis.RankingKey, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error()), utils.ErrFromRedis)
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
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), food), utils.ErrFromRedis)
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
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error()), utils.ErrFromMysqlDB)
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
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), redisKey, foodName, score), utils.ErrFromRedis)
	}

	return nil
}

func (g *RankingFoodRepository) PreviousRankingExist(ctx context.Context) (int, error) {
	// Check if previous ranking exists
	previousExist, err := _redis.Client.Exists(ctx, _redis.PrevRankingKey).Result()
	if err != nil {
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error()), utils.ErrFromRedis)
	}

	return int(previousExist), nil
}

func (g *RankingFoodRepository) FindOneFoods(ctx context.Context, foodID int) (string, error) {
	// Find food name by food ID
	var foodName string
	fmt.Println(foodID)
	err := g.GormDB.WithContext(ctx).Model(&mysql.Foods{}).Select("name").Where("id = ?", foodID).First(&foodName).Error
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodID), utils.ErrFromMysqlDB)
	}

	return foodName, nil
}

func (g *RankingFoodRepository) ExpireRanking(ctx context.Context, key string) error {
	// Set expiration time for key
	err := g.RedisClient.Expire(ctx, key, 30*time.Minute).Err() // RedisClient는 Redis 연결을 나타내는 변수입니다.
	if err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), key), utils.ErrFromRedis)
	}

	return nil
}
