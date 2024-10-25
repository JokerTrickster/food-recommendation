package handler

import (
	"context"
	"encoding/json"
	"fmt"
	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	_redis "main/utils/db/redis"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

type MetaFoodHandler struct {
	UseCase _interface.IMetaFoodUseCase
}

func NewMetaFoodHandler(c *echo.Echo, useCase _interface.IMetaFoodUseCase) _interface.IMetaFoodHandler {
	handler := &MetaFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/meta", handler.Meta)
	return handler
}

// 메타 데이터 가져오기
// @Router /v0.1/foods/meta [get]
// @Summary 메타 데이터 가져오기
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
// @Success 200 {object} response.ResMetaData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *MetaFoodHandler) Meta(c echo.Context) error {
	ctx := context.Background()
	// 캐시 체크
	//business logic
	cacheKey := "meta:category"
	metaData, err := _redis.Client.Get(ctx, cacheKey).Result()
	if metaData == "" {
		res, err := d.UseCase.Meta(ctx)
		if err != nil {
			return err
		}
		fmt.Println(res)
		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return err
		}
		err = _redis.Client.Set(ctx, cacheKey, data, 1*time.Hour).Err()
		if err != nil {
			return err
		}

		// 캐시 히트 여부
		c.Response().Header().Set("X-Cache-Hit", "false")

		// 4. DB에서 조회한 데이터 반환
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		return err
	}
	var res response.ResMetaData
	if err := json.Unmarshal([]byte(metaData), &res); err != nil {
		return err
	}
	// 캐시 히트 여부
	c.Response().Header().Set("X-Cache-Hit", "true")

	return c.JSON(http.StatusOK, res)
}
