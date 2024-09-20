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

func NewRecommendFoodRepository(gormDB *gorm.DB) _interface.IRecommendFoodRepository {
	return &RecommendFoodRepository{GormDB: gormDB}
}
func (d *RecommendFoodRepository) FindOneUser(ctx context.Context, uID uint) (*mysql.Users, error) {

	user := mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", uID).First(&user)
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), _errors.ErrUserNotFound.Error(), utils.ErrFromClient)
	}
	return &user, nil
}

func (d *RecommendFoodRepository) SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) (*mysql.Foods, error) {
	foods := mysql.Foods{}
	// 존재 여부 확인
	err := d.GormDB.WithContext(ctx).Model(&foods).Where("name = ? AND time_id = ? AND type_id = ? AND scenario_id = ? and theme_id = ? and flavor_id = ?", foodDTO.Name, foodDTO.TimeID, foodDTO.TypeID, foodDTO.ScenarioID, foodDTO.ThemeID, foodDTO.FlavorID).First(&foods).Error

	if err == nil {
		// 데이터가 이미 존재함
		return &foods, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 데이터베이스 오류
		return &mysql.Foods{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}

	// 데이터가 존재하지 않으므로 저장
	if err := d.GormDB.WithContext(ctx).Create(&foodDTO).Error; err != nil {
		return &mysql.Foods{}, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return foodDTO, nil
}
