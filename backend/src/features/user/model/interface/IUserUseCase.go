package _interface

import (
	"context"
	"main/features/user/model/entity"
	"main/features/user/model/response"
)

type IGetUserUseCase interface {
	Get(c context.Context, uID uint) (response.ResGetUser, error)
}

type IUpdateUserUseCase interface {
	Update(c context.Context, entity *entity.UpdateUserEntity) error
}
