package repository

import (
	"context"
	_errors "main/features/user/model/errors"
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewGetUserRepository(gormDB *gorm.DB) _interface.IGetUserRepository {
	return &GetUserRepository{GormDB: gormDB}
}
func (d *GetUserRepository) FindOneUser(ctx context.Context, uID uint) (*mysql.Users, error) {

	user := mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", uID).First(&user)
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), utils.HandleError(_errors.ErrUserNotFound.Error(),uID), utils.ErrFromClient)
	}
	return &user, nil
}
