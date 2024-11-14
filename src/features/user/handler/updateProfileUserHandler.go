package handler

import (
	"main/features/user/model/entity"
	_interface "main/features/user/model/interface"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateProfileUserHandler struct {
	UseCase _interface.IUpdateProfileUserUseCase
}

func NewUpdateProfileUserHandler(c *echo.Echo, useCase _interface.IUpdateProfileUserUseCase) _interface.IUpdateProfileUserHandler {
	handler := &UpdateProfileUserHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/users/profiles/image", handler.UpdateProfile, mw.TokenChecker)
	return handler
}

// 유저 프로필 이미지 저장하기
// @Router /v0.1/users/profiles/image [post]
// @Summary 유저 프로필 이미지 저장하기
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
// @Param tkn header string true "accessToken"
// @Param image formData file false "프로필 이미지 파일"
// @Produce json
// @Success 200 {object} response.ResUpdateProfileUser
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags user
func (d *UpdateProfileUserHandler) UpdateProfile(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	e := &entity.UpdateProfileUserEntity{
		UserID: uID,
		Image:  file,
	}
	res, err := d.UseCase.UpdateProfile(ctx, e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
