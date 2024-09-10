package handler

import (
	"main/features/food/repository"
	"main/features/food/usecase"
	"main/utils/db/mysql"
	_redis "main/utils/db/redis"

	"github.com/labstack/echo/v4"
)

func NewFoodHandler(c *echo.Echo) {
	NewRecommendFoodHandler(c, usecase.NewRecommendFoodUseCase(repository.NewRecommendFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewSelectFoodHandler(c, usecase.NewSelectFoodUseCase(repository.NewSelectFoodRepository(mysql.GormMysqlDB, _redis.Client), mysql.DBTimeOut))
	NewHistoryFoodHandler(c, usecase.NewHistoryFoodUseCase(repository.NewHistoryFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewMetaFoodHandler(c, usecase.NewMetaFoodUseCase(repository.NewMetaFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
	NewRankingFoodHandler(c, usecase.NewRankingFoodUseCase(repository.NewRankingFoodRepository(mysql.GormMysqlDB, _redis.Client), mysql.DBTimeOut))
	NewImageUploadFoodHandler(c, usecase.NewImageUploadFoodUseCase(repository.NewImageUploadFoodRepository(mysql.GormMysqlDB), mysql.DBTimeOut))
}
