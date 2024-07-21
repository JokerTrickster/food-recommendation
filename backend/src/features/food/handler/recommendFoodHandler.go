package handler

import (
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/features/food/model/request"
	"main/features/food/model/response"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecommendFoodHandler struct {
	UseCase _interface.IRecommendFoodUseCase
}

func NewRecommendFoodHandler(c *echo.Echo, useCase _interface.IRecommendFoodUseCase) _interface.IRecommendFoodHandler {
	handler := &RecommendFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods/recommend", handler.Recommend, mw.TokenChecker)
	return handler
}

// 음식 추천 받기
// @Router /v0.1/foods/recommend [post]
// @Summary 음식 추천 받기
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
// @Param type body request.ReqRecommendFood true "type"
// @Produce json
// @Success 200 {object} response.ResRecommendFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *RecommendFoodHandler) Recommend(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	req := &request.ReqRecommendFood{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	//business logic
	entity := entity.RecommendFoodEntity{
		Type:     req.Type,
		Scenario: req.Scenario,
		Time:     req.Time,
		UserID:   uID,
	}
	if req.PreviousAnswer != "" {
		entity.PreviousAnswer = req.PreviousAnswer
	}

	data, err := d.UseCase.Recommend(ctx, entity)
	if err != nil {
		return err
	}

	res := response.ResRecommendFood{
		FoodNames: data,
	}

	return c.JSON(http.StatusOK, res)
}
