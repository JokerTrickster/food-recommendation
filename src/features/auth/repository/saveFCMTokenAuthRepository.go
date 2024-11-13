package repository

import (
	"context"
	"errors"
	_interface "main/features/auth/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewSaveFCMTokenAuthRepository(gormDB *gorm.DB) _interface.ISaveFCMTokenAuthRepository {
	return &SaveFCMTokenAuthRepository{GormDB: gormDB}
}

func (d *SaveFCMTokenAuthRepository) SaveFCMToken(ctx context.Context, userTokenDTO *mysql.UserTokens) error {
	// 기존에 userID에 해당하는 토큰이 존재하면 업데이트하고 없다면 새로 생성
	var existingToken mysql.UserTokens
	err := d.GormDB.Where("user_id = ?", userTokenDTO.UserID).First(&existingToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 기존 레코드가 없으면 새로 생성
			err = d.GormDB.Create(&userTokenDTO).Error
			if err != nil {
				return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), userTokenDTO), utils.ErrFromMysqlDB)
			}
		} else {
			return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), userTokenDTO), utils.ErrFromMysqlDB)
		}
	} else {
		// 기존 레코드가 있으면 업데이트
		existingToken.Token = userTokenDTO.Token
		err = d.GormDB.Save(&existingToken).Error
		if err != nil {
			return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), userTokenDTO), utils.ErrFromMysqlDB)
		}
	}
	return nil
}
