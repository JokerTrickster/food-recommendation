package repository

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewNaverOauthAuthRepository(gormDB *gorm.DB) _interface.INaverOauthAuthRepository {
	return &NaverOauthAuthRepository{GormDB: gormDB}
}

func (g *NaverOauthAuthRepository) DeleteToken(ctx context.Context, uID uint) error {
	token := mysql.Tokens{
		UserID: uID,
	}
	result := g.GormDB.Model(&token).Where("user_id = ?", uID).Delete(&token)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), utils.HandleError(result.Error.Error(),uID), utils.ErrFromInternal)
	}
	return nil
}
func (g *NaverOauthAuthRepository) SaveToken(ctx context.Context, uID uint, accessToken, refreshToken string, refreshTknExpiredAt int64) error {
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

func (g *NaverOauthAuthRepository) InsertOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error) {
	result := g.GormDB.WithContext(ctx).Create(&user)
	if result.RowsAffected == 0 {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("failed user insert",user), utils.ErrFromMysqlDB)
	}
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(),user), utils.ErrFromMysqlDB)
	}
	return user, nil
}

func (g *NaverOauthAuthRepository) FindOneUser(ctx context.Context, user *mysql.Users) (*mysql.Users, error) {
	var newUser mysql.Users
	result := g.GormDB.WithContext(ctx).Where("email = ?", user.Email).First(&newUser)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(),user), utils.ErrFromMysqlDB)
	}
	return &newUser, nil
}
