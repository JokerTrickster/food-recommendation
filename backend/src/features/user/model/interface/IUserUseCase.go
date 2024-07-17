package _interface

import (
	"context"
	"main/features/user/model/entity"
	"main/utils/db/mysql"
)

type IGetUserUseCase interface {
	Get(c context.Context, uID uint) (*mysql.Users, error)
}

type IUpdateUserUseCase interface {
	Update(c context.Context, entity *entity.UpdateUserEntity) error
}
