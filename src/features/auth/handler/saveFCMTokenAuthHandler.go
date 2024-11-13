package handler

import (
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SaveFCMTokenAuthHandler struct {
	UseCase _interface.ISaveFCMTokenAuthUseCase
}

func NewSaveFCMTokenAuthHandler(c *echo.Echo, useCase _interface.ISaveFCMTokenAuthUseCase) _interface.ISaveFCMTokenAuthHandler {
	handler := &SaveFCMTokenAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/fcm/token", handler.SaveFCMToken, mw.TokenChecker)
	return handler
}

// fcm 토큰 저장
// @Router /v0.1/auth/fcm/token [post]
// @Summary fcm 토큰 저장
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description INVALID_EMAIL_OR_PASSWORD : 비밀번호 또는 이메일 잘못 요청
// @Description ■ errCode with 401
// @Description INVALID_AUTH_CODE : 인증 코드 검증 실패
// @Description TOKEN_BAD : 잘못된 토큰
// @Description INVALID_ACCESS_TOKEN : 잘못된 액세스 토큰
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Description PLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패
// @Param tkn header string true "accessToken"
// @Param json body request.ReqSaveFCMToken true "fcm 토큰"
// @Produce json
// @Success 200 {object} boolean
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *SaveFCMTokenAuthHandler) SaveFCMToken(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	req := &request.ReqSaveFCMToken{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	err := d.UseCase.SaveFCMToken(ctx, uID, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
