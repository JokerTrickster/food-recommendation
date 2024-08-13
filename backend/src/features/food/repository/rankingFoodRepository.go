package repository

import (
	"context"
	_interface "main/features/food/model/interface"
	_redis "main/utils/db/redis"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRankingFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IRankingFoodRepository {
	return &RankingFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *RankingFoodRepository) FindAllRanking(ctx context.Context) ([]string, error) {
	foodList, err := g.RedisClient.ZRevRange(ctx, _redis.RankingKey, 0, 9).Result()
	if err != nil {
		return nil, err
	}
	return foodList, nil
}
