package usecase

import (
	"context"
	"fmt"
	_interface "main/features/user/model/interface"
	"main/utils/db/mysql"
	"time"
)

type GetUserUseCase struct {
	Repository     _interface.IGetUserRepository
	ContextTimeout time.Duration
}

func NewGetUserUseCase(repo _interface.IGetUserRepository, timeout time.Duration) _interface.IGetUserUseCase {
	return &GetUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GetUserUseCase) Get(c context.Context, uID uint) (*mysql.Users, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Print(ctx)
	// 유저 정보를 가져온다.
	userDTO, err := d.Repository.FindOneUser(ctx, uID)
	if err != nil {
		return nil, err
	}
	if userDTO.Sex == "" {
		return nil, nil
	}
	return userDTO, nil
}
