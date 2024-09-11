package handler

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type V02GoogleOauthAuthHandler struct {
	UseCase _interface.IV02GoogleOauthAuthUseCase
}

func NewV02GoogleOauthAuthHandler(c *echo.Echo, useCase _interface.IV02GoogleOauthAuthUseCase) _interface.IV02GoogleOauthAuthHandler {
	handler := &V02GoogleOauthAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.2/auth/google", handler.V02GoogleOauth)
	return handler
}

// [app] google oauth 로그인
// @Router /v0.2/auth/google [post]
// @Summary [app] google oauth 로그인
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param json body request.ReqV02GoogleOauth true "구글 토큰"
// @Produce json
// @Success 200 {object} response.ResV02GoogleOauth
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *V02GoogleOauthAuthHandler) V02GoogleOauth(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqV02GoogleOauth{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	res, err := d.UseCase.V02GoogleOauth(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
