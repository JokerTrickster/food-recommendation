package handler

import (
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/features/food/model/request"

	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type V12RecommendFoodHandler struct {
	UseCase _interface.IV12RecommendFoodUseCase
}

func NewV12RecommendFoodHandler(c *echo.Echo, useCase _interface.IV12RecommendFoodUseCase) _interface.IV12RecommendFoodHandler {
	handler := &V12RecommendFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v1.2/foods/recommend", handler.V12Recommend, mw.TokenChecker)
	return handler
}

// 음식 추천 받기 (llm 제거)
// @Router /v1.2/foods/recommend [post]
// @Summary 음식 추천 받기 (llm 제거)
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
// @Param tkn header string true "accessToken"
// @Param type body request.ReqV12RecommendFood true "type"
// @Produce json
// @Success 200 {object} response.ResV12RecommendFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *V12RecommendFoodHandler) V12Recommend(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	req := &request.ReqV12RecommendFood{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	//business logic
	entity := entity.V12RecommendFoodEntity{
		Types:     req.Types,
		Scenarios: req.Scenarios,
		Times:     req.Times,
		Themes:    req.Themes,
		UserID:    uID,
	}
	if req.PreviousAnswer != "" {
		entity.PreviousAnswer = req.PreviousAnswer
	}

	res, err := d.UseCase.V12Recommend(ctx, entity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
