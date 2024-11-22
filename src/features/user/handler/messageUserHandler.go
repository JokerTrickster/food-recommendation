package handler

import (
	"context"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"

	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MessageUserHandler struct {
	UseCase _interface.IMessageUserUseCase
}

func NewMessageUserHandler(c *echo.Echo, useCase _interface.IMessageUserUseCase) _interface.IMessageUserHandler {
	handler := &MessageUserHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/users/message", handler.Message)
	return handler
}

// 유저 메시지 전송하기
// @Router /v0.1/users/message [post]
// @Summary 유저 메시지 전송하기
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저가 존재하지 않음
// @Description ■ errCode with 401
// @Description INVALID_AUTH_CODE : 인증 코드 검증 실패
// @Description TOKEN_BAD : 잘못된 토큰
// @Description INVALID_ACCESS_TOKEN : 잘못된 액세스 토큰
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Description PLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패
// @param json body request.ReqMessageUser true "수정한 데이터만 전달"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *MessageUserHandler) Message(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqMessageUser{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	err := d.UseCase.Message(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
