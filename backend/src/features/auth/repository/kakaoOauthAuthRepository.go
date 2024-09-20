package repository

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewKakaoOauthAuthRepository(gormDB *gorm.DB) _interface.IKakaoOauthAuthRepository {
	return &KakaoOauthAuthRepository{GormDB: gormDB}
}

func (g *KakaoOauthAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), result.Error.Error(), utils.ErrFromInternal)
	}
	return nil
}
func (g *KakaoOauthAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
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

func (g *KakaoOauthAuthRepository) InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), "failed user insert", utils.ErrFromMysqlDB)
	}
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), result.Error.Error(), utils.ErrFromMysqlDB)
	}
	return user, nil
}

func (g *KakaoOauthAuthRepository) FindOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error) {
	var newUser mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", user.Email).First(&newUser)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), result.Error.Error(), utils.ErrFromMysqlDB)
	}
	return &newUser, nil
}