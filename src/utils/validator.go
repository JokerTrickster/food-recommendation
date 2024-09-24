package utils

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// validator instance
var val = validator.New()

// ValidateStruct : validate struct
func ValidateStruct(class interface{}) error {
	if err := val.Struct(class); err != nil {
		return ErrorMsg(context.TODO(), ErrBadParameter, Trace(), err.Error(), ErrFromClient)
	}
	return nil
}

// ValidateReq : validate REST API req body
func ValidateReq(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return ErrorMsg(context.TODO(), ErrBadParameter, Trace(), err.Error(), ErrFromClient)
	}
	if err := val.Struct(req); err != nil {
		return ErrorMsg(context.TODO(), ErrBadParameter, Trace(), err.Error(), ErrFromClient)
	}
	return nil
}

// ValidateRes : validate REST API res body
func ValidateRes(c echo.Context, res interface{}) error {
	if err := val.Struct(res); err != nil {
		return ErrorMsg(context.TODO(), ErrBadParameter, Trace(), err.Error(), ErrFromClient)
	}
	return nil
}
