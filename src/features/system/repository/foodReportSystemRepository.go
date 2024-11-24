package repository

import (
	"context"
	_interface "main/features/system/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewFoodReportSystemRepository(gormDB *gorm.DB) _interface.IFoodReportSystemRepository {
	return &FoodReportSystemRepository{GormDB: gormDB}
}

func (d *FoodReportSystemRepository) FindFoodWithoutImage(ctx context.Context) ([]*mysql.FoodImages, error) {
	var foodImages []*mysql.FoodImages
	result := d.GormDB.Where("image = ?", "food_default.png").Find(&foodImages)
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(), foodImages), utils.ErrFromMysqlDB)
	}
	return foodImages, nil
}

func (d *FoodReportSystemRepository) FindFoodWithoutNutrient(ctx context.Context) ([]*mysql.Foods, error) {

	var foodsWithoutNutrients []*mysql.Foods

	query := `
        SELECT f.*
        FROM foods f
        LEFT JOIN nutrients n ON f.name = n.food_name
        WHERE n.food_name IS NULL AND f.deleted_at IS NULL
    `

	if err := d.GormDB.WithContext(ctx).Raw(query).Scan(&foodsWithoutNutrients).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodsWithoutNutrients), utils.ErrFromMysqlDB)
	}

	return foodsWithoutNutrients, nil

}
