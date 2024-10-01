package handler

import (
	_interface "main/features/system/model/interface"
	"main/features/system/model/request"
	mw "main/middleware"
	"main/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportSystemHandler struct {
	UseCase _interface.IReportSystemUseCase
}

func NewReportSystemHandler(c *echo.Echo, useCase _interface.IReportSystemUseCase) _interface.IReportSystemHandler {
	handler := &ReportSystemHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/system/report", handler.Report, mw.TokenChecker)
	return handler
}

// 1:1 문의하기
// @Router /v0.1/system/report [post]
// @Summary 1:1 문의하기
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
// @Param json body request.ReqReport true "json body"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags system
func (d *ReportSystemHandler) Report(c echo.Context) error {
	ctx, uID, _ := utils.CtxGenerate(c)
	req := &request.ReqReport{}
	if err := utils.ValidateReq(c, req); err != nil {
		return err
	}
	err := d.UseCase.Report(ctx, uID, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}
