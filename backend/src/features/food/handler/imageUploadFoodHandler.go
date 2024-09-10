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
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Description PLAYER_STATE_CHANGE_FAILED : 플레이어 상태 변경 실패
// @Param foodID formData string false "food ID"
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

	foodID, err := strconv.Atoi(c.FormValue("foodID"))
	if err != nil {
		return err
	}
	entity := entity.ImageUploadFoodEntity{
		FoodID: foodID,
		Image:  file,
	}

	//business logic
	err = d.UseCase.ImageUpload(ctx, entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
