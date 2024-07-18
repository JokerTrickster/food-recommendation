package repository

import (
	"context"
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
