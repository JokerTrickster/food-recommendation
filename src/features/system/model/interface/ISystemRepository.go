package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IReportSystemRepository interface {
	SaveReport(ctx context.Context, reportDTO *mysql.Reports) error
}
