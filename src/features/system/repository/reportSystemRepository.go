package repository

import (
	"context"
	_interface "main/features/system/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewReportSystemRepository(gormDB *gorm.DB) _interface.IReportSystemRepository {
	return &ReportSystemRepository{GormDB: gormDB}
}

func (d *ReportSystemRepository) SaveReport(ctx context.Context, reportDTO *mysql.Reports) error {

	result := d.GormDB.Create(&reportDTO)
	if result.Error != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(), reportDTO), utils.ErrFromMysqlDB)
	}
	return nil
}
