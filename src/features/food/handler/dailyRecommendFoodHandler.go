package handler

import (
	"context"
	"encoding/json"
	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	_redis "main/utils/db/redis"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
)

type DailyRecommendFoodHandler struct {
	UseCase _interface.IDailyRecommendFoodUseCase
}

func NewDailyRecommendFoodHandler(c *echo.Echo, useCase _interface.IDailyRecommendFoodUseCase) _interface.IDailyRecommendFoodHandler {
	handler := &DailyRecommendFoodHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/foods/daily-recommend", handler.DailyRecommend)
	return handler
}

// 오늘의 음식 추천 받기 (3개)
// @Router /v0.1/foods/daily-recommend [get]
// @Summary 오늘의 음식 추천 받기 (3개))
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
// @Produce json
// @Success 200 {object} response.ResDailyRecommendFood
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags food
func (d *DailyRecommendFoodHandler) DailyRecommend(c echo.Context) error {
	ctx := context.Background()
	//business logic
	// Redis 캐시 처리

	foodData, err := _redis.Client.Get(ctx, _redis.DailyKey).Result()
	if foodData == "" {
		// 2. 캐시에 데이터가 없을 경우 UseCase에서 조회
		res, err := d.UseCase.DailyRecommend(ctx)
		if err != nil {
			return err
		}

		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return err
		}
		err = _redis.Client.Set(ctx, _redis.DailyKey, data, 1*time.Hour).Err()
		if err != nil {
			return err
		}

		// 캐시 히트 여부
		c.Response().Header().Set("X-Cache-Hit", "false")

		// 4. DB에서 조회한 데이터 반환
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		// Redis 오류 처리
		return err
	}

	// 5. 캐시된 데이터가 있을 경우 반환
	var res response.ResDailyRecommendFood
	if err := json.Unmarshal([]byte(foodData), &res); err != nil {
		return err
	}

	// 캐시 히트 여부
	c.Response().Header().Set("X-Cache-Hit", "true")

	return c.JSON(http.StatusOK, res)
}
