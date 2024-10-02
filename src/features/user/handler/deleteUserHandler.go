package handler

import (
	_interface "main/features/user/model/interface"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeleteUserHandler struct {
	UseCase _interface.IDeleteUserUseCase
}

func NewDeleteUserHandler(c *echo.Echo, useCase _interface.IDeleteUserUseCase) _interface.IDeleteUserHandler {
	handler := &DeleteUserHandler{
		UseCase: useCase,
	}
	c.DELETE("/v0.1/users", handler.Delete, mw.TokenChecker)
	return handler
}

// 회원 탈퇴하기
// @Router /v0.1/users [delete]
// @Summary 회원 탈퇴하기
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
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *DeleteUserHandler) Delete(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)

	err := d.UseCase.Delete(ctx, uID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
