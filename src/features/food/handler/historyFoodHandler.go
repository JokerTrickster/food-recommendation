package handler

import (
	_interface "main/features/food/model/interface"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HistoryFoodHandler struct {
	UseCase _interface.IHistoryFoodUseCase
}

func NewHistoryFoodHandler(c *echo.Echo, useCase _interface.IHistoryFoodUseCase) _interface.IHistoryFoodHandler {
	handler := &HistoryFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/history", handler.History, mw.TokenChecker)
	return handler
}

// 음식 히스토리 가쟈오기
// @Router /v0.1/foods/history [get]
// @Summary 음식 히스토리 가쟈오기
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
// @Param tkn header string true "accessToken"
// @Produce json
// @Success 200 {object} response.ResHistoryFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *HistoryFoodHandler) History(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)

	//business logic

	res, err := d.UseCase.History(ctx, uID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
