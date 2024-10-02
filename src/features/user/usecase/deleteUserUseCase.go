package usecase

import (
	"context"
	_interface "main/features/user/model/interface"
	"main/utils"
	"time"
)

type DeleteUserUseCase struct {
	Repository     _interface.IDeleteUserRepository
	ContextTimeout time.Duration
}

func NewDeleteUserUseCase(repo _interface.IDeleteUserRepository, timeout time.Duration) _interface.IDeleteUserUseCase {
	return &DeleteUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *DeleteUserUseCase) Delete(c context.Context, uID uint) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	if uID == 1 {
		return utils.ErrorMsg(ctx, utils.ErrBadParameter, utils.Trace(), utils.HandleError("탈퇴할 수 없습니다.", uID), utils.ErrFromClient)
	}
	err := d.Repository.DeleteUser(ctx, uID)
	if err != nil {
		return err
	}

	return nil
}
