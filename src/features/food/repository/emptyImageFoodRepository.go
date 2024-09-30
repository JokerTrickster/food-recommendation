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

func (g *EmptyImageFoodRepository) FindAllEmptyImageFoods(ctx context.Context) ([]mysql.FoodImages, error) {
	var foodImages []mysql.FoodImages
	if err := g.GormDB.WithContext(ctx).Where("image = ?", "food_default.png").Find(&foodImages).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error()), utils.ErrFromMysqlDB)
	}
	return foodImages, nil
}
