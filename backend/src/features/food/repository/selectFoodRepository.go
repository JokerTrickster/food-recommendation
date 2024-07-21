package repository

import (
	"context"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewSelectFoodRepository(gormDB *gorm.DB) _interface.ISelectFoodRepository {
	return &SelectFoodRepository{GormDB: gormDB}
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
