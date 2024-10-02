package _interface

import "github.com/labstack/echo/v4"

type IGetUserHandler interface {
	Get(c echo.Context) error
}

type IUpdateUserHandler interface {
	Update(c echo.Context) error
}

type IDeleteUserHandler interface {
	Delete(c echo.Context) error
}
