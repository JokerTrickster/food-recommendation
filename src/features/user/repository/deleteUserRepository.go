package repository

import (
	"context"
	_errors "main/features/user/model/errors"
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewDeleteUserRepository(gormDB *gorm.DB) _interface.IDeleteUserRepository {
	return &DeleteUserRepository{GormDB: gormDB}
}

func (d *DeleteUserRepository) DeleteUser(ctx context.Context, uID uint) error {
	result := d.GormDB.Where("id = ?", uID).Delete(&mysql.Users{})
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(), uID), utils.ErrFromInternal)
	}
	if result.RowsAffected == 0 {
		return utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), utils.HandleError(_errors.ErrUserNotFound.Error(), uID), utils.ErrFromClient)
	}
	return nil
}
