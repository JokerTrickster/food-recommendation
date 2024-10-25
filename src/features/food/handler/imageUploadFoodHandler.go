package handler

import (
	"context"
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageUploadFoodHandler struct {
	UseCase _interface.IImageUploadFoodUseCase
}

func NewImageUploadFoodHandler(c *echo.Echo, useCase _interface.IImageUploadFoodUseCase) _interface.IImageUploadFoodHandler {
	handler := &ImageUploadFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods/image", handler.ImageUpload)
	return handler
}

// 음식 이미지 업로드하기
// @Router /v0.1/foods/image [post]
// @Summary 음식 이미지 업로드하기
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
// @Param foodImageID formData string false "food image ID"
// @Param image formData file false "음식 이미지 파일"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *ImageUploadFoodHandler) ImageUpload(c echo.Context) error {
	ctx := context.Background()
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	foodImageID, err := strconv.Atoi(c.FormValue("foodImageID"))
	if err != nil {
		return err
	}
	entity := entity.ImageUploadFoodEntity{
		FoodID: foodImageID,
		Image:  file,
	}

	//business logic
	err = d.UseCase.ImageUpload(ctx, entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
