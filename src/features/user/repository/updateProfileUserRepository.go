package repository

import (
	"context"
	_errors "main/features/user/model/errors"
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewUpdateProfileUserRepository(gormDB *gorm.DB) _interface.IUpdateProfileUserRepository {
	return &UpdateProfileUserRepository{GormDB: gormDB}
}

func (d *UpdateProfileUserRepository) UpdateProfileImage(ctx context.Context, userID uint, filename string) error {
	user := &mysql.Users{}
	result := d.GormDB.Model(&user).Where("id = ?", userID).Update("image", filename)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(), user), utils.ErrFromInternal)
	}
	if result.RowsAffected == 0 {
		return utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), utils.HandleError(_errors.ErrUserNotFound.Error(), user), utils.ErrFromClient)
	}
	return nil
}
