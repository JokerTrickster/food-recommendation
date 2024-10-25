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

type SelectFoodHandler struct {
	UseCase _interface.ISelectFoodUseCase
}

func NewSelectFoodHandler(c *echo.Echo, useCase _interface.ISelectFoodUseCase) _interface.ISelectFoodHandler {
	handler := &SelectFoodHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/foods/select", handler.Select, mw.TokenChecker)
	return handler
}

// 음식 선택하기
// @Router /v0.1/foods/select [post]
// @Summary 음식 선택하기
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
// @Param type body request.ReqSelectFood true "type"
// @Produce json
// @Success 200 {object} response.ResSelectFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *SelectFoodHandler) Select(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	req := &request.ReqSelectFood{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}

	//business logic
	e := entity.SelectFoodEntity{
		Types:     req.Types,
		Times:     req.Times,
		Name:      req.Name,
		Themes:    req.Themes,
		Flavors:   req.Flavors,
		Scenarios: req.Scenarios,
		UserID:    uID,
	}

	res, err := d.UseCase.Select(ctx, e)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
