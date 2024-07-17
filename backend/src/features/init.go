package features

import (
	"net/http"

	authHandler "main/features/auth/handler"
	userHandler "main/features/user/handler"

	"github.com/labstack/echo/v4"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	authHandler.NewAuthHandler(e)
	userHandler.NewUserHandler(e)

	return nil
}
