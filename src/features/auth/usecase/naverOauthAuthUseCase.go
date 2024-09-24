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

type NaverOauthAuthUseCase struct {
	Repository     _interface.INaverOauthAuthRepository
	ContextTimeout time.Duration
}

func NewNaverOauthAuthUseCase(repo _interface.INaverOauthAuthRepository, timeout time.Duration) _interface.INaverOauthAuthUseCase {
	return &NaverOauthAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *NaverOauthAuthUseCase) NaverOauth(c context.Context, req *request.ReqNaverOauth) (response.ResNaverOauth, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 토큰 검증
	oauthData, err := utils.NaverValidate(ctx, req.Token)
	if err != nil {
		return response.ResNaverOauth{}, err
	}

	// 유저 생성
	userDTO := CreateGoogleUserDTO(oauthData)
	var newUserDTO *mysql.Users
	//유저 존재 체크
	newUserDTO, err = d.Repository.FindOneUser(ctx, userDTO)
	if err != nil {
		return response.ResNaverOauth{}, err
	}
	if newUserDTO == nil {
		// 유저 정보 insert
		newUserDTO, err = d.Repository.InsertOneUser(ctx, userDTO)
		if err != nil {
			return response.ResNaverOauth{}, err
		}
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := utils.GenerateToken(newUserDTO.Email, newUserDTO.ID)
	if err != nil {
		return response.ResNaverOauth{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, newUserDTO.ID)
	if err != nil {
		return response.ResNaverOauth{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, newUserDTO.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResNaverOauth{}, err
	}
	res := response.ResNaverOauth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       newUserDTO.ID,
	}

	return res, nil
}
