package usecase

import (
	"context"
	"fmt"
	_interface "main/features/auth/model/interface"
	"main/features/auth/model/request"
	"main/features/auth/model/response"
	"main/utils"
	"main/utils/db/mysql"
	"time"
)

type KakaoOauthAuthUseCase struct {
	Repository     _interface.IKakaoOauthAuthRepository
	ContextTimeout time.Duration
}

func NewKakaoOauthAuthUseCase(repo _interface.IKakaoOauthAuthRepository, timeout time.Duration) _interface.IKakaoOauthAuthUseCase {
	return &KakaoOauthAuthUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *KakaoOauthAuthUseCase) KakaoOauth(c context.Context, req *request.ReqKakaoOauth) (response.ResKakaoOauth, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 토큰 검증
	oauthData, err := utils.KakaoValidate(ctx, req.Token)
	if err != nil {
		fmt.Println(err)
		return response.ResKakaoOauth{}, err
	}

	// 유저 생성
	userDTO := CreateGoogleUserDTO(oauthData)
	var newUserDTO *mysql.Users
	//유저 존재 체크
	newUserDTO, err = d.Repository.FindOneUser(ctx, userDTO)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	if newUserDTO == nil {
		// 유저 정보 insert
		newUserDTO, err = d.Repository.InsertOneUser(ctx, userDTO)
		if err != nil {
			return response.ResKakaoOauth{}, err
		}
	}

	// token create
	accessToken, _, refreshToken, refreshTknExpiredAt, err := utils.GenerateToken(newUserDTO.Email, newUserDTO.ID)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}

	// 기존 토큰 제거
	err = d.Repository.DeleteToken(ctx, newUserDTO.ID)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	// token db save
	err = d.Repository.SaveToken(ctx, newUserDTO.ID, accessToken, refreshToken, refreshTknExpiredAt)
	if err != nil {
		return response.ResKakaoOauth{}, err
	}
	res := response.ResKakaoOauth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       newUserDTO.ID,
	}

	return res, nil
}
