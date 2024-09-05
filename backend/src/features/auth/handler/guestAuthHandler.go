package handler

import (
	"context"
	"encoding/json"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/response"
	_redis "main/utils/db/redis"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type GuestAuthHandler struct {
	UseCase _interface.IGuestAuthUseCase
}

func NewGuestAuthHandler(c *echo.Echo, useCase _interface.IGuestAuthUseCase) _interface.IGuestAuthHandler {
	handler := &GuestAuthHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/auth/guest", handler.Guest)
	return handler
}

// 게스트 로그인
// @Router /v0.1/auth/guest [post]
// @Summary 게스트 로그인
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_EXIST : 유저가 존재하지 않음
// @Description USER_ALREADY_EXISTED : 유저가 이미 존재
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} response.ResGuest
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags auth
func (d *GuestAuthHandler) Guest(c echo.Context) error {
	ctx := context.Background()
	// 레디스 캐시 처리
	cacheKey := "guest_token"
	guestAuth, err := _redis.Client.Get(ctx, cacheKey).Result()
	if guestAuth == "" {
		// 2. 캐시에 데이터가 없을 경우 DB에서 조회
		res, err := d.UseCase.Guest(ctx)
		if err != nil {
			return err
		}

		// 3. 조회된 데이터를 Redis에 캐시 (예: 1시간 TTL)
		data, err := json.Marshal(res)
		if err != nil {
			return err
		}
		err = _redis.Client.Set(ctx, cacheKey, data, time.Hour).Err()
		if err != nil {
			return err
		}
		//캐시 히트 여부
		c.Response().Header().Set("X-Cache-Hit", "false")
		// 4. DB에서 조회한 데이터 반환
		return c.JSON(http.StatusOK, res)
	} else if err != nil {
		// Redis 오류 처리
		return err
	}

	var res response.ResGuest
	if err := json.Unmarshal([]byte(guestAuth), &res); err != nil {
		return err
	}
	c.Response().Header().Set("X-Cache-Hit", "true")
	return c.JSON(http.StatusOK, res)
}
