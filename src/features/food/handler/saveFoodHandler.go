package handler

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/features/food/model/request"

	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SaveFoodHandler struct {
	UseCase _interface.ISaveFoodUseCase
}

func NewSaveFoodHandler(c *echo.Echo, useCase _interface.ISaveFoodUseCase) _interface.ISaveFoodHandler {
	handler := &SaveFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods", handler.Save)
	return handler
}

// 음식 이름 저장하기
// @Router /v0.1/foods [post]
// @Summary 음식 이름 저장하기
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
// @Param type body request.ReqSaveFood true "이름, 시간, 상황, 맛, 테마, 종류"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *SaveFoodHandler) Save(c echo.Context) error {
	req := &request.ReqSaveFood{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	err := d.UseCase.Save(context.TODO(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
