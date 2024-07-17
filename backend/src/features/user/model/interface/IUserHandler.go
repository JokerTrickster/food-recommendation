package _interface

import "github.com/labstack/echo/v4"

type IGetUserHandler interface {
	Get(c echo.Context) error
}
