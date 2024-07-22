package handler

import (
	"main/features/food/repository"
	"main/features/food/usecase"
	"main/utils/db/mysql"

	"github.com/labstack/echo/v4"
)

func NewFoodHandler(c *echo.Echo) {
	NewRecommendFoodHandler(c, usecase.NewRecommendFoodUseCase(repository.NewRecommendFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewSelectFoodHandler(c, usecase.NewSelectFoodUseCase(repository.NewSelectFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewHistoryFoodHandler(c, usecase.NewHistoryFoodUseCase(repository.NewHistoryFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
}
