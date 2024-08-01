package handler

import (
	"context"
	_interface "main/features/auth/model/interface"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GuestAuthHandler struct {
	UseCase _interface.IGuestAuthUseCase
}

func NewGuestAuthHandler(c *echo.Echo, useCase _interface.IGuestAuthUseCase) _interface.IGuestAuthHandler {
	handler := &GuestAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/guest", handler.Guest)
	return handler
}

// 게스트 로그인
// @Router /v0.1/auth/guest [post]
// @Summary 게스트 로그인
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} response.ResGuest
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *GuestAuthHandler) Guest(c echo.Context) error {
	ctx := context.Background()

	res, err := d.UseCase.Guest(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
