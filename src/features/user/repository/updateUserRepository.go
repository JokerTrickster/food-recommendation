package repository

import (
	"context"
	_errors "main/features/user/model/errors"
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewUpdateUserRepository(gormDB *gorm.DB) _interface.IUpdateUserRepository {
	return &UpdateUserRepository{GormDB: gormDB}
}
func (d *UpdateUserRepository) FindOneAndUpdateUser(ctx context.Context, userDTO *mysql.Users) error {

	result := d.GormDB.Model(&userDTO).Where("id = ?", userDTO.ID).Updates(&userDTO)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(),userDTO), utils.ErrFromInternal)
	}
	if result.RowsAffected == 0 {
		return utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), utils.HandleError(_errors.ErrUserNotFound.Error(),userDTO), utils.ErrFromClient)
	}
	return nil
}
