package handler

import (
	"main/utils"
	"net/http"

	_interface "main/features/food/model/interface"

	"github.com/labstack/echo/v4"
)

type RankingFoodHandler struct {
	UseCase _interface.IRankingFoodUseCase
}

func NewRankingFoodHandler(c *echo.Echo, useCase _interface.IRankingFoodUseCase) _interface.IRankingFoodHandler {
	handler := &RankingFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/rank", handler.Ranking)
	return handler
}

// 실시간 음식 랭킹 가져오기 (TOP 10)
// @Router /v0.1/foods/rank [get]
// @Summary 실시간 음식 랭킹 가져오기 (TOP 10)
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
// @Produce json
// @Success 200 {object} response.ResRankingFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *RankingFoodHandler) Ranking(c echo.Context) error {
	ctx, _, _ := utils.CtxGenerate(c)

	//business logic

	res, err := d.UseCase.Ranking(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
