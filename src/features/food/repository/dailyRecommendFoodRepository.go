package repository

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewDailyRecommendFoodRepository(gormDB *gorm.DB) _interface.IDailyRecommendFoodRepository {
	return &DailyRecommendFoodRepository{GormDB: gormDB}
}
func (d *DailyRecommendFoodRepository) FindOneFood(ctx context.Context, foodName string) (*mysql.Foods, error) {
	food := mysql.Foods{}
	if err := d.GormDB.WithContext(ctx).Model(&food).Where("name = ?", foodName).First(&food).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodName), utils.ErrFromMysqlDB)
	}
	return &food, nil

}
func (d *DailyRecommendFoodRepository) FindOneFoodImage(ctx context.Context, foodID int) (string, error) {
	foodImage := mysql.FoodImages{}
	if err := d.GormDB.WithContext(ctx).Model(&foodImage).Where("id = ?", foodID).First(&foodImage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "food_default.png", nil
		}
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodID), utils.ErrFromMysqlDB)
	}
	return foodImage.Image, nil
}
