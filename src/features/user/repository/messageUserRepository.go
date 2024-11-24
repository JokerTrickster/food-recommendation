package repository

import (
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func NewMessageUserRepository(gormDB *gorm.DB) _interface.IMessageUserRepository {
	return &MessageUserRepository{GormDB: gormDB}
}

func (d *MessageUserRepository) FindOnePushToken(ctx context.Context, uID uint) (string, error) {
	var userToken *mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", uID).First(&userToken).Error
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("error finding push token", err), utils.ErrFromMysqlDB)
	}
	return userToken.Token, nil
}

func (d *MessageUserRepository) FindOneAlarm(ctx context.Context, uID uint) (bool, error) {
	var user *mysql.Users
	err := d.GormDB.Where("id = ?", uID).First(&user).Error
	if err != nil {
		return false, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("error finding user", err), utils.ErrFromMysqlDB)
	}
	return *user.Push, nil
}
