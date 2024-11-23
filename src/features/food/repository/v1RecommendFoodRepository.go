package repository

import (
	"context"
	"errors"
	"fmt"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewV1RecommendFoodRepository(gormDB *gorm.DB) _interface.IV1RecommendFoodRepository {
	return &V1RecommendFoodRepository{GormDB: gormDB}
}

func (d *V1RecommendFoodRepository) FindOneV1RecommendFood(ctx context.Context, query string) (*mysql.Foods, error) {
	food := mysql.Foods{} // 포인터가 아닌 구조체로 초기화
	result := d.GormDB.WithContext(ctx).Raw(query).Scan(&food)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no records found for the given query")
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no records found for the given query")
	}

	return &food, nil // 반환할 때 포인터로 반환
}

func (d *V1RecommendFoodRepository) SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) (*mysql.Foods, error) {
	foods := mysql.Foods{}
	// 존재 여부 확인
	err := d.GormDB.WithContext(ctx).Model(&foods).Where("name = ? AND time_id = ? AND type_id = ? AND scenario_id = ? and theme_id = ? and flavor_id = ?", foodDTO.Name, foodDTO.TimeID, foodDTO.TypeID, foodDTO.ScenarioID, foodDTO.ThemeID, foodDTO.FlavorID).First(&foods).Error

	if err == nil {
		// 데이터가 이미 존재함
		return &foods, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 데이터베이스 오류
		return &mysql.Foods{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodDTO), utils.ErrFromMysqlDB)
	}

	// 데이터가 존재하지 않으므로 저장
	if err := d.GormDB.WithContext(ctx).Create(&foodDTO).Error; err != nil {
		return &mysql.Foods{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodDTO), utils.ErrFromMysqlDB)
	}
	return foodDTO, nil
}
func (d *V1RecommendFoodRepository) FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *mysql.FoodImages) (*mysql.FoodImages, error) {
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

func (d *V1RecommendFoodRepository) CountV1RecommendFood(ctx context.Context, query string) (int, error) {
	var foods []mysql.Foods
	// Count를 Raw 쿼리와 함께 사용하는 대신, GORM의 쿼리 빌더를 사용
	err := d.GormDB.WithContext(ctx).Raw(query).Scan(&foods).Error
	if err != nil {
		return 0, err
	}
	return len(foods), nil
}
func (d *V1RecommendFoodRepository) FindOneFoodImage(ctx context.Context, id int) (string, error) {
	foodImage := mysql.FoodImages{}
	err := d.GormDB.WithContext(ctx).Where("id = ?", id).First(&foodImage).Error
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), id), utils.ErrFromMysqlDB)
	}
	return foodImage.Image, nil
}

func (d *V1RecommendFoodRepository) FindOneNutrient(ctx context.Context, foodName string) (*mysql.Nutrients, error) {
	nutrient := mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", foodName).First(&nutrient).Error
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), foodName), utils.ErrFromMysqlDB)
	}
	return &nutrient, nil
}

func (d *V1RecommendFoodRepository) FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) (*mysql.Nutrients, error) {
	nutrient := mysql.Nutrients{}
	err := d.GormDB.WithContext(ctx).Where("food_name = ?", nutrientDTO.FoodName).First(&nutrient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 데이터가 없을 경우 저장
			if err := d.GormDB.WithContext(ctx).Create(&nutrientDTO).Error; err != nil {
				return nil, err
			}
			return nutrientDTO, nil
		}
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(_errors.ErrServerError.Error()+err.Error(), nutrientDTO), utils.ErrFromMysqlDB)
	}
	return &nutrient, nil
}
