package repository

import (
	"context"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewHistoryFoodRepository(gormDB *gorm.DB) _interface.IHistoryFoodRepository {
	return &HistoryFoodRepository{GormDB: gormDB}
}

func (g *HistoryFoodRepository) FindOneFood(ctx context.Context, foodID uint) (*mysql.Foods, error) {
	food := mysql.Foods{}
	if err := g.GormDB.WithContext(ctx).Model(&food).Where("id = ?", foodID).First(&food).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return &food, nil
}

func (g *HistoryFoodRepository) FindAllFoodHistory(ctx context.Context, userID uint) ([]mysql.FoodHistory, error) {
	foodHistoryList := []mysql.FoodHistory{}
	if err := g.GormDB.WithContext(ctx).Where("user_id = ?", userID).Find(&foodHistoryList).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return foodHistoryList, nil
}
