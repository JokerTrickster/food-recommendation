package repository

import (
	"context"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"
	_redis "main/utils/db/redis"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewSelectFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.ISelectFoodRepository {
	return &SelectFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *SelectFoodRepository) FindOneFood(ctx context.Context, foodDTO *mysql.Foods) (uint, error) {
	food := mysql.Foods{}
	if err := g.GormDB.WithContext(ctx).Model(&food).Where(foodDTO).First(&food).Error; err != nil {
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return food.ID, nil
}
func (g *SelectFoodRepository) InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *mysql.FoodHistory) error {
	if err := g.GormDB.WithContext(ctx).Create(&foodHistoryDTO).Error; err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}

	return nil
}

func (g *SelectFoodRepository) IncrementFoodRanking(ctx context.Context, foodID string, score float64) error {
	redisKey := _redis.RankingKey
	_, err := g.RedisClient.ZAdd(ctx, redisKey, redis.Z{Score: score, Member: foodID}).Result()
	if err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromRedis)
	}
	return nil
}
