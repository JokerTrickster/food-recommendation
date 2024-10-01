package handler

import (
	"main/features/system/repository"
	"main/features/system/usecase"
	"main/utils/db/mysql"

	"github.com/labstack/echo/v4"
)

func NewSystemHandler(c *echo.Echo) {
	NewReportSystemHandler(c, usecase.NewReportSystemUseCase(repository.NewReportSystemRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
}
