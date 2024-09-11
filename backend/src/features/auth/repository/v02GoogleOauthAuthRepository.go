package repository

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewV02GoogleOauthAuthRepository(gormDB *gorm.DB) _interface.IV02GoogleOauthAuthRepository {
	return &V02GoogleOauthAuthRepository{GormDB: gormDB}
}

func (g *V02GoogleOauthAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), result.Error.Error(), utils.ErrFromInternal)
	}
	return nil
}
func (g *V02GoogleOauthAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
	token := mysql.Tokens{
		UserID:           uID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		RefreshExpiredAt: refreshTknExpiredAt,
	}
	result := g.GormDB.Model(&token).Create(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), result.Error.Error(), utils.ErrFromInternal)
	}
	return nil
}

func (g *V02GoogleOauthAuthRepository) InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), "failed user insert", utils.ErrFromMysqlDB)
	}
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), result.Error.Error(), utils.ErrFromMysqlDB)
	}
	return user, nil
}
