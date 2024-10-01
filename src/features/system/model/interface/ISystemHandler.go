package _interface

import "github.com/labstack/echo/v4"

type IReportSystemHandler interface {
	Report(c echo.Context) error
}
