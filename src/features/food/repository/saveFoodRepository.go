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

func NewSaveFoodRepository(gormDB *gorm.DB) _interface.ISaveFoodRepository {
	return &SaveFoodRepository{GormDB: gormDB}
}

func (d *SaveFoodRepository) SaveFood(ctx context.Context, foodDTO *mysql.Foods) (uint, error) {
	foods := mysql.Foods{}
	// 존재 여부 확인
	err := d.GormDB.WithContext(ctx).Model(&foods).Where("name = ?  ", foodDTO.Name).First(&foods).Error

	if err == nil {
		// 데이터가 이미 존재함
		return foods.ID, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 데이터베이스 오류
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodDTO), utils.ErrFromMysqlDB)
	}

	// 데이터가 존재하지 않으므로 저장
	if err := d.GormDB.WithContext(ctx).Create(&foodDTO).Error; err != nil {
		return 0, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodDTO), utils.ErrFromMysqlDB)
	}
	return foodDTO.ID, nil
}

func (d *SaveFoodRepository) FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *mysql.FoodImages) (*mysql.FoodImages, error) {
	foodImage := mysql.FoodImages{}

	// food_name 기준으로 데이터 조회
	if err := d.GormDB.WithContext(ctx).Where("name = ?", foodImageDTO.Name).First(&foodImage).Error; err != nil {
		// 데이터가 없을 경우 ErrRecordNotFound 발생
		if err == gorm.ErrRecordNotFound {
			// 데이터를 저장
			if err := d.GormDB.WithContext(ctx).Create(&foodImageDTO).Error; err != nil {
				return &mysql.FoodImages{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodImageDTO), utils.ErrFromMysqlDB)
			}
			// 저장된 데이터를 반환
			return foodImageDTO, nil
		}
		// 다른 에러 처리
		return &mysql.FoodImages{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodImageDTO), utils.ErrFromMysqlDB)
	}

	// 조회된 데이터를 반환
	return &foodImage, nil
}

func (d *SaveFoodRepository) SaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) error {
	// 데이터 저장
	if err := d.GormDB.WithContext(ctx).Create(&nutrientDTO).Error; err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), utils.ErrFromMysqlDB)
	}
	return nil
}

func (d *SaveFoodRepository) FindCategoryIDs(ctx context.Context, categories []string) ([]uint, error) {
	categoryIDs := make([]uint, 0)
	// 카테고리 조회
	result := d.GormDB.WithContext(ctx).Model(&mysql.Categories{}).Where("name IN ?", categories).Pluck("id", &categoryIDs)
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+result.Error.Error(), categories), utils.ErrFromMysqlDB)
	}
	return categoryIDs, nil
}

func (d *SaveFoodRepository) SaveFoodCategory(ctx context.Context, foodID uint, categoryIDs []uint) error {
	// 음식 카테고리 저장
	for _, categoryID := range categoryIDs {
		foodCategoryDTO := mysql.FoodCategories{FoodID: int(foodID), CategoryID: int(categoryID)}
		if err := d.GormDB.WithContext(ctx).Create(&foodCategoryDTO).Error; err != nil {
			return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodCategoryDTO), utils.ErrFromMysqlDB)
		}
	}
	return nil
}
