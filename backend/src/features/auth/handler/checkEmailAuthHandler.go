package handler

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckEmailAuthHandler struct {
	UseCase _interface.ICheckEmailAuthUseCase
}

func NewCheckEmailAuthHandler(c *echo.Echo, useCase _interface.ICheckEmailAuthUseCase) _interface.ICheckEmailAuthHandler {
	handler := &CheckEmailAuthHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/auth/email/check", handler.CheckEmail)
	return handler
}

// 이메일 중복 체크
// @Router /v0.1/auth/email/check [get]
// @Summary 이메일 중복 체크
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Description PLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패
// @Param email query string true "email"
// @Produce json
// @Success 200 {object} boolean
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *CheckEmailAuthHandler) CheckEmail(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqCheckEmail{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	err := d.UseCase.CheckEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
