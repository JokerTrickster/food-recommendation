package usecase

import (
	"context"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	"main/features/auth/model/response"
	"main/utils"
	"main/utils/db/mysql"
	"time"
)

type V02GoogleOauthAuthUseCase struct {
	Repository     _interface.IV02GoogleOauthAuthRepository
	ContextTimeout time.Duration
}

func NewV02GoogleOauthAuthUseCase(repo _interface.IV02GoogleOauthAuthRepository, timeout time.Duration) _interface.IV02GoogleOauthAuthUseCase {
	return &V02GoogleOauthAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *V02GoogleOauthAuthUseCase) V02GoogleOauth(c context.Context, req *request.ReqV02GoogleOauth) (response.ResV02GoogleOauth, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 토큰 검증
	oauthData, err := utils.GoogleValidate(ctx, req.Token)
	if err != nil {
		return response.ResV02GoogleOauth{}, err
	}

	// 유저 생성
	userDTO := CreateGoogleUserDTO(oauthData)
	var newUserDTO *mysql.Users
	//유저 존재 체크
	newUserDTO, err = d.Repository.FindOneUser(ctx, userDTO)
	if err != nil {
		return response.ResV02GoogleOauth{}, err
	}
	if newUserDTO == nil {
		// 유저 정보 insert
		newUserDTO, err = d.Repository.InsertOneUser(ctx, userDTO)
		if err != nil {
			return response.ResV02GoogleOauth{}, err
		}
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := utils.GenerateToken(newUserDTO.Email, newUserDTO.ID)
	if err != nil {
		return response.ResV02GoogleOauth{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, newUserDTO.ID)
	if err != nil {
		return response.ResV02GoogleOauth{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, newUserDTO.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResV02GoogleOauth{}, err
	}
	res := response.ResV02GoogleOauth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       newUserDTO.ID,
	}

	return res, nil
}
