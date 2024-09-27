package repository

import (
	"context"
	_errors "main/features/auth/model/errors"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewSignupAuthRepository(gormDB *gorm.DB) _interface.ISignupAuthRepository {
	return &SignupAuthRepository{GormDB: gormDB}
}

func (g *SignupAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(),uID), utils.ErrFromInternal)
	}
	return nil
}
func (g *SignupAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(),uID), utils.ErrFromInternal)
	}
	return nil
}
func (g *SignupAuthRepository) UserCheckByEmail(ctx context.Context, email string) error {
	var user mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		return nil
	} else {
		return utils.ErrorMsg(ctx, utils.ErrUserAlreadyExisted, utils.Trace(), utils.HandleError(_errors.ErrUserAlreadyExisted.Error(),email), utils.ErrFromClient)
	}
}
func (g *SignupAuthRepository) InsertOneUser(ctx context.Context, user mysql.Users) error {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("failed user insert",user), utils.ErrFromMysqlDB)
	}
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(),user), utils.ErrFromMysqlDB)
	}
	return nil
}
