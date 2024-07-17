package handler

import (
	_interface "main/features/user/model/interface"
	"main/features/user/model/response"

	_errors "main/features/user/model/errors"
	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetUserHandler struct {
	UseCase _interface.IGetUserUseCase
}

func NewGetUserHandler(c *echo.Echo, useCase _interface.IGetUserUseCase) _interface.IGetUserHandler {
	handler := &GetUserHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/users/check", handler.Get, mw.TokenChecker)
	return handler
}

// 유저 프로필 가져오기
// @Router /v0.1/users/check [get]
// @Summary 유저 프로필 가져오기
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
// @Success 200 {object} response.ResGetUser
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *GetUserHandler) Get(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)

	userDTO, err := d.UseCase.Get(ctx, uID)
	if err != nil {
		return err
	}
	if userDTO == nil {
		return utils.ErrorMsg(ctx, utils.ErrProfileNotFount, utils.Trace(), _errors.ErrProfileNotFound.Error(), utils.ErrFromClient)
	}
	res := response.ResGetUser{
		Birth: userDTO.Birth,
		Sex:   userDTO.Sex,
	}
	return c.JSON(http.StatusOK, res)
}
