package handler

import (
	"context"
	_interface "main/features/system/model/interface"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodReportSystemHandler struct {
	UseCase _interface.IFoodReportSystemUseCase
}

func NewFoodReportSystemHandler(c *echo.Echo, useCase _interface.IFoodReportSystemUseCase) _interface.IFoodReportSystemHandler {
	handler := &FoodReportSystemHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/system/foods/report", handler.FoodReport)
	return handler
}

// 음식 이미지, 영양 성분 없는 음식 리포트
// @Router /v0.1/system/foods/report [post]
// @Summary 음식 이미지, 영양 성분 없는 음식 리포트
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
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags system
func (d *FoodReportSystemHandler) FoodReport(c echo.Context) error {
	ctx := context.Background()
	err := d.UseCase.FoodReport(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
