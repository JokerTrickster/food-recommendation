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
