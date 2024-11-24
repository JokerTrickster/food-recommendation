package repository

import (
	"context"
	_interface "main/features/user/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewAllMessageUserRepository(gormDB *gorm.DB) _interface.IAllMessageUserRepository {
	return &AllMessageUserRepository{GormDB: gormDB}
}

func (d *AllMessageUserRepository) FindUsersForNotifications(ctx context.Context) ([]*mysql.Users, error) {
	users := []*mysql.Users{}
	err := d.GormDB.WithContext(ctx).Where("push = ?", true).Find(&users).Error
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("error finding users for notifications", err), utils.ErrFromMysqlDB)
	}
	return users, nil
}

func (d *AllMessageUserRepository) FindOnePushToken(ctx context.Context, uID uint) (string, error) {
	var userToken *mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", uID).First(&userToken).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError("error finding push token", err), utils.ErrFromMysqlDB)
	}
	return userToken.Token, nil
}
