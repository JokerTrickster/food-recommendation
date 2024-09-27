package repository

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewGuestAuthRepository(gormDB *gorm.DB) _interface.IGuestAuthRepository {
	return &GuestAuthRepository{GormDB: gormDB}
}

func (g *GuestAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(),uID,accessToken,refreshToken), utils.ErrFromInternal)
	}
	return nil
}

func (g *GuestAuthRepository) FindOneAndUpdateUser(ctx context.Context, email, password string) (mysql.Users, error) {
	user := mysql.Users{
		Email: email,
	}

	err := g.GormDB.WithContext(ctx).Where("email = ? and password = ? and provider = ?", email, password, "test").First(&user).Error
	if err != nil {
		return mysql.Users{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(err.Error(),email,password), utils.ErrFromInternal)
	}
	return user, nil
}
