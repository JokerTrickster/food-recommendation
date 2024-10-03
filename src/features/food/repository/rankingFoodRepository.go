package repository

import (
	"context"
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/utils"
	_redis "main/utils/db/redis"

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
			Name:  z.Member.(string),
			Score: z.Score,
		}
		result = append(result, rankFood)
	}

	return result, nil
}

func (g *RankingFoodRepository) FindRankingFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error) {
	// gorm에서 food_histories 테이블에서 top 10가져오기
	var results []struct {
		name  string
		Count int64
	}
	// SQL 쿼리 실행
	err := g.GormDB.Table("food_histories").
		Select("name, COUNT(name) as count").
		Group("name").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error()), utils.ErrFromMysqlDB)
	}

	// 결과에서 음식 이름 추출
	topFoods := make([]*entity.RankFoodRedis, 0)
	for _, r := range results {
		topFoods = append(topFoods, &entity.RankFoodRedis{
			Name:  r.name,
			Score: float64(r.Count),
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
