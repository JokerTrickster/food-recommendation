package repository

import (
	"context"
	_errors "main/features/auth/model/errors"
	_interface "main/features/auth/model/interface"
	"main/utils"

	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewCheckEmailAuthRepository(gormDB *gorm.DB) _interface.ICheckEmailAuthRepository {
	return &CheckEmailAuthRepository{GormDB: gormDB}
}

func (g *CheckEmailAuthRepository) CheckEmail(ctx context.Context, email string) error {
	user := mysql.Users{
		Email: email,
	}
	//이메일 중복 체크
	result := g.GormDB.WithContext(ctx).Model(&user).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), _errors.ErrUserNotFound.Error(), utils.ErrFromClient)
	}
	if result.RowsAffected == 0 {
		return nil
	} else {
		return utils.ErrorMsg(ctx, utils.ErrUserAlreadyExisted, utils.Trace(), _errors.ErrUserAlreadyExisted.Error(), utils.ErrFromClient)
	}
}
