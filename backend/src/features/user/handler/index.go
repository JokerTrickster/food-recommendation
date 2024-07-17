package handler

import (
	"main/features/user/repository"
	"main/features/user/usecase"
	"main/utils/db/mysql"

	"github.com/labstack/echo/v4"
)

func NewUserHandler(c *echo.Echo) {
	NewGetUserHandler(c, usecase.NewGetUserUseCase(repository.NewGetUserRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewUpdateUserHandler(c, usecase.NewUpdateUserUseCase(repository.NewUpdateUserRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
}
