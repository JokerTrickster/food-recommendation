package _errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrUserAlreadyExisted = errors.New("user already existed")
	ErrCodeNotFound       = errors.New("code not found")
	ErrInvalidAuthCode    = errors.New("invalid auth code")
)
