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
	if result.Error != nil && result.Error.Error() != gorm.ErrRecordNotFound.Error() {
		return utils.ErrorMsg(ctx, utils.ErrUserNotFound, utils.Trace(), utils.HandleError(_errors.ErrUserNotFound.Error()+result.Error.Error(), email), utils.ErrFromClient)
	}
	if result.RowsAffected == 1 {
		return utils.ErrorMsg(ctx, utils.ErrUserAlreadyExisted, utils.Trace(), utils.HandleError(_errors.ErrUserAlreadyExisted.Error(), email), utils.ErrFromClient)
	}
	return nil
}
