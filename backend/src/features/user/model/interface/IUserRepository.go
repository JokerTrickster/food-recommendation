package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IGetUserRepository interface {
	FindOneUser(ctx context.Context, uID uint) (*mysql.Users, error)
}
