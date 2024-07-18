package _interface

import "github.com/labstack/echo/v4"

type IRecommendFoodHandler interface {
	Recommend(c echo.Context) error
}
