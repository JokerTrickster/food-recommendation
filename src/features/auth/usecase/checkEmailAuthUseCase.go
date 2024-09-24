package usecase

import (
	"context"
	"fmt"
	_interface "main/features/auth/model/interface"
	"time"
)

type CheckEmailAuthUseCase struct {
	Repository     _interface.ICheckEmailAuthRepository
	ContextTimeout time.Duration
}

func NewCheckEmailAuthUseCase(repo _interface.ICheckEmailAuthRepository, timeout time.Duration) _interface.ICheckEmailAuthUseCase {
	return &CheckEmailAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *CheckEmailAuthUseCase) CheckEmail(c context.Context, email string) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Println(ctx)
	// email check
	err := d.Repository.CheckEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
