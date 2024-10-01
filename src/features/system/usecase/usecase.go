package usecase

import (
	"main/features/system/model/request"
	"main/utils/db/mysql"
)

func CreateReportDTO(userID uint, req *request.ReqReport) *mysql.Reports {
	return &mysql.Reports{
		UserID: int(userID),
		Reason: req.Reason,
	}
}
