package usecase

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/response"
	"main/utils"
	"time"
)

type GuestAuthUseCase struct {
	Repository     _interface.IGuestAuthRepository
	ContextTimeout time.Duration
}

func NewGuestAuthUseCase(repo _interface.IGuestAuthRepository, timeout time.Duration) _interface.IGuestAuthUseCase {
	return &GuestAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *GuestAuthUseCase) Guest(c context.Context) (response.ResGuest, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// user check
	user, err := d.Repository.FindOneAndUpdateUser(ctx, "test@test.com", "asdasd123")
	if err != nil {
		return response.ResGuest{}, err
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		return response.ResGuest{}, err
	}

	// token db save
	err = d.Repository.SaveToken(ctx, user.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResGuest{}, err
	}

	//response create
	res := response.ResGuest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}
