package handler

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KakaoOauthAuthHandler struct {
	UseCase _interface.IKakaoOauthAuthUseCase
}

func NewKakaoOauthAuthHandler(c *echo.Echo, useCase _interface.IKakaoOauthAuthUseCase) _interface.IKakaoOauthAuthHandler {
	handler := &KakaoOauthAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/kakao", handler.KakaoOauth)
	return handler
}

// [app] kakao oauth 로그인
// @Router /v0.1/auth/kakao [post]
// @Summary [app] kakao oauth 로그인
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param json body request.ReqKakaoOauth true "구글 토큰"
// @Produce json
// @Success 200 {object} response.ResKakaoOauth
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *KakaoOauthAuthHandler) KakaoOauth(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqKakaoOauth{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	res, err := d.UseCase.KakaoOauth(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
