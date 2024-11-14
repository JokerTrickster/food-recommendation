package usecase

import (
	"context"
	"main/features/user/model/entity"
	_interface "main/features/user/model/interface"
	"main/features/user/model/response"
	"main/utils/aws"
	"time"
)

type UpdateProfileUserUseCase struct {
	Repository     _interface.IUpdateProfileUserRepository
	ContextTimeout time.Duration
}

func NewUpdateProfileUserUseCase(repo _interface.IUpdateProfileUserRepository, timeout time.Duration) _interface.IUpdateProfileUserUseCase {
	return &UpdateProfileUserUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *UpdateProfileUserUseCase) UpdateProfile(c context.Context, e *entity.UpdateProfileUserEntity) (response.ResUpdateProfileUser, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	//s3 이미지 업로드 한다.
	filename := aws.FileNameGenerateRandom()
	err := aws.ImageUpload(ctx, e.Image, filename, aws.ImgTypeProfile)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}

	//유저 정보를 업데이트 한다.
	err = d.Repository.UpdateProfileImage(ctx, e.UserID, filename)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}

	//s3 이미지 url을 응답한다.
	url, err := aws.ImageGetSignedURL(ctx, filename, aws.ImgTypeProfile)
	if err != nil {
		return response.ResUpdateProfileUser{}, err
	}
	res := response.ResUpdateProfileUser{
		Image: url,
	}

	return res, nil
}
