package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IGetUserUseCase interface {
	Get(c context.Context, uID uint) (*mysql.Users, error)
}
