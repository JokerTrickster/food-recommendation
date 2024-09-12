package handler

import (
	"context"
	_interface "main/features/food/model/interface"

	"net/http"

	"github.com/labstack/echo/v4"
)

type DailyRecommendFoodHandler struct {
	UseCase _interface.IDailyRecommendFoodUseCase
}

func NewDailyRecommendFoodHandler(c *echo.Echo, useCase _interface.IDailyRecommendFoodUseCase) _interface.IDailyRecommendFoodHandler {
	handler := &DailyRecommendFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/daily-recommend", handler.DailyRecommend)
	return handler
}

// 오늘의 음식 추천 받기 (3개)
// @Router /v0.1/foods/daily-recommend [get]
// @Summary 오늘의 음식 추천 받기 (3개))
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Description PLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패
// @Produce json
// @Success 200 {object} response.ResDailyRecommendFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *DailyRecommendFoodHandler) DailyRecommend(c echo.Context) error {
	ctx := context.Background()
	//business logic

	res, err := d.UseCase.DailyRecommend(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
