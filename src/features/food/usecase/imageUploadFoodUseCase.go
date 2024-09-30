package usecase

import (
	"context"
	"fmt"

	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	_aws "main/utils/aws"

	"time"
)

type ImageUploadFoodUseCase struct {
	Repository     _interface.IImageUploadFoodRepository
	ContextTimeout time.Duration
}

func NewImageUploadFoodUseCase(repo _interface.IImageUploadFoodRepository, timeout time.Duration) _interface.IImageUploadFoodUseCase {
	return &ImageUploadFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *ImageUploadFoodUseCase) ImageUpload(c context.Context, e entity.ImageUploadFoodEntity) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 랜덤으로 이미지 이름 생성
	filename := fmt.Sprintf("%s.png", _aws.FileNameGenerateRandom())
	// 디비에 이미지 파일 이름 저장
	err := d.Repository.FindOneAndUpdateFoodImages(ctx, uint(e.FoodID), filename)
	if err != nil {
		return err
	}
	// s3 이미지 파일 업로드
	err = _aws.ImageUpload(ctx, e.Image, filename, _aws.ImgTypeFood)
	if err != nil {
		return err
	}

	return nil
}
