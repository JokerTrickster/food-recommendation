package _interface

import (
	"context"
	"main/features/system/model/request"
)

type IReportSystemUseCase interface {
	Report(c context.Context, uID uint, req *request.ReqReport) error
}
