package handler

import (
	"main/features/user/model/entity"
	_interface "main/features/user/model/interface"
	"main/features/user/model/request"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateUserHandler struct {
	UseCase _interface.IUpdateUserUseCase
}

func NewUpdateUserHandler(c *echo.Echo, useCase _interface.IUpdateUserUseCase) _interface.IUpdateUserHandler {
	handler := &UpdateUserHandler{
		UseCase: useCase,
	}
	c.PUT("/v0.1/users/profile", handler.Update, mw.TokenChecker)
	return handler
}

// 유저 프로필 저장하기
// @Router /v0.1/users/profile [put]
// @Summary 유저 프로필 저장하기
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
// @param json body request.ReqUpdateUser true "모든 데이터 전달 요청(수정하는 데이터만 보내면 에러 발생)"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *UpdateUserHandler) Update(c echo.Context) error {
	ctx, uID, email := utils.CtxGenerate(c)
	req := &request.ReqUpdateUser{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	entity := entity.UpdateUserEntity{
		UserID: uID,
		Email:  email,
		Birth:  req.Birth,
		Sex:    req.Sex,
		Name:   req.Name,
	}
	err := d.UseCase.Update(ctx, &entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
