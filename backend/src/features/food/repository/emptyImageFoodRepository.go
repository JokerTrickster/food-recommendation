package repository

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewEmptyImageFoodRepository(gormDB *gorm.DB) _interface.IEmptyImageFoodRepository {
	return &EmptyImageFoodRepository{GormDB: gormDB}
}

func (g *EmptyImageFoodRepository) FindAllEmptyImageFoods(ctx context.Context) ([]mysql.Foods, error) {
	var foods []mysql.Foods
	if err := g.GormDB.WithContext(ctx).Where("image = ?", "food_default.png").Find(&foods).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromMysqlDB)
	}
	return foods, nil
}
