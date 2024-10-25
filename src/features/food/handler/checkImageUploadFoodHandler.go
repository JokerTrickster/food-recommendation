package handler

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/features/food/model/request"
	"main/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckImageUploadFoodHandler struct {
	UseCase _interface.ICheckImageUploadFoodUseCase
}

func NewCheckImageUploadFoodHandler(c *echo.Echo, useCase _interface.ICheckImageUploadFoodUseCase) _interface.ICheckImageUploadFoodHandler {
	handler := &CheckImageUploadFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods/images/upload-check", handler.CheckImageUpload)
	return handler
}

// 음식 이미지 업로드 체크하기
// @Router /v0.1/foods/images/upload-check [post]
// @Summary 음식 이미지 업로드 체크하기
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
// @Description GEMINI_INTERNAL_SERVER : Gemini 서버 내부 오류
// @Param type body request.ReqCheckImageUploadFood true "type"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *CheckImageUploadFoodHandler) CheckImageUpload(c echo.Context) error {
	ctx := context.Background()
	req := &request.ReqCheckImageUploadFood{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	//business logic
	err := d.UseCase.CheckImageUpload(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
