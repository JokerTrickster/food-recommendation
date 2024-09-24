package handler

import (
	"context"
	_interface "main/features/food/model/interface"

	"net/http"

	"github.com/labstack/echo/v4"
)

type EmptyImageFoodHandler struct {
	UseCase _interface.IEmptyImageFoodUseCase
}

func NewEmptyImageFoodHandler(c *echo.Echo, useCase _interface.IEmptyImageFoodUseCase) _interface.IEmptyImageFoodHandler {
	handler := &EmptyImageFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/empty-image", handler.EmptyImage)
	return handler
}

// 음식 이미지 없는 음식 정보 가져오기
// @Router /v0.1/foods/empty-image [get]
// @Summary 음식 이미지 없는 음식 정보 가져오기
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
// @Success 200 {object} response.ResEmptyImageFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *EmptyImageFoodHandler) EmptyImage(c echo.Context) error {
	ctx := context.Background()

	//business logic
	res, err := d.UseCase.EmptyImage(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
