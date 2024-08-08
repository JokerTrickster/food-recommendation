package _interface

import "github.com/labstack/echo/v4"

type ISignupAuthHandler interface {
	Signup(c echo.Context) error
}

type ISigninAuthHandler interface {
	Signin(c echo.Context) error
}

type ILogoutAuthHandler interface {
	Logout(c echo.Context) error
}

type IReissueAuthHandler interface {
	Reissue(c echo.Context) error
}

type IRequestPasswordAuthHandler interface {
	RequestPassword(c echo.Context) error
}
type IValidatePasswordAuthHandler interface {
	ValidatePassword(c echo.Context) error
}

type ICheckEmailAuthHandler interface {
	CheckEmail(c echo.Context) error
}

type IGuestAuthHandler interface {
	Guest(c echo.Context) error
}

type IGoogleOauthAuthHandler interface {
	GoogleOauth(c echo.Context) error
}
type IGoogleOauthCallbackAuthHandler interface {
	GoogleOauthCallback(c echo.Context) error
}

type IV02GoogleOauthCallbackAuthHandler interface {
	V02GoogleOauthCallback(c echo.Context) error
}
