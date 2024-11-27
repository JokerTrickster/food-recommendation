package repository

import (
	"context"
	"errors"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewV12RecommendFoodRepository(gormDB *gorm.DB) _interface.IV12RecommendFoodRepository {
	return &V12RecommendFoodRepository{GormDB: gormDB}
}

func (d *V12RecommendFoodRepository) FindOneV12RecommendFood(ctx context.Context, query string) (*mysql.Foods, error) {
	food := mysql.Foods{} // 포인터가 아닌 구조체로 초기화
	result := d.GormDB.WithContext(ctx).Raw(query).Scan(&food)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorMsg(ctx, utils.ErrFoodNotFound, utils.Trace(), "no matching record found", utils.ErrFromClient)
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, utils.ErrorMsg(ctx, utils.ErrFoodNotFound, utils.Trace(), "no matching record found", utils.ErrFromClient)
	}

	return &food, nil // 반환할 때 포인터로 반환
}

func (d *V12RecommendFoodRepository) FindOneFoodImage(ctx context.Context, id int) (string, error) {
	foodImage := mysql.FoodImages{}
	err := d.GormDB.WithContext(ctx).Where("id = ?", id).First(&foodImage).Error
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), id), utils.ErrFromMysqlDB)
	}
	return foodImage.Image, nil
}

func (d *V12RecommendFoodRepository) FindOneNutrient(ctx context.Context, foodName string) (*mysql.Nutrients, error) {
	nutrient := mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", foodName).First(&nutrient).Error
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodName), utils.ErrFromMysqlDB)
	}
	return &nutrient, nil
}

func (d *V12RecommendFoodRepository) FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) (*mysql.Nutrients, error) {
	nutrient := mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", nutrientDTO.FoodName).First(&nutrient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 데이터가 없을 경우 저장
			if err := d.GormDB.WithContext(ctx).Create(&nutrientDTO).Error; err != nil {
				return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), utils.ErrFromMysqlDB)
			}
			return nutrientDTO, nil
		}
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), utils.ErrFromMysqlDB)
	}
	return &nutrient, nil
}
