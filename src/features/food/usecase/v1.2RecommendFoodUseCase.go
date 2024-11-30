package usecase

import (
	"context"
	"main/features/food/model/entity"
	"main/features/food/model/response"
	"main/utils/aws"

	_interface "main/features/food/model/interface"
	"time"
)

type V12RecommendFoodUseCase struct {
	Repository     _interface.IV12RecommendFoodRepository
	ContextTimeout time.Duration
}

func NewV12RecommendFoodUseCase(repo _interface.IV12RecommendFoodRepository, timeout time.Duration) _interface.IV12RecommendFoodUseCase {
	return &V12RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}
func (d *V12RecommendFoodUseCase) V12Recommend(c context.Context, e entity.V12RecommendFoodEntity) (response.ResV12RecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	query := CreateV12RecommendQuery(e)

	query += " ORDER BY RAND() LIMIT 1"
	food, err := d.Repository.FindOneV12RecommendFood(ctx, query)
	if err != nil {
		return response.ResV12RecommendFood{}, err
	}
	//food image ID로 이미지 URL을 가져온다.
	image, err := d.Repository.FindOneFoodImage(ctx, food.ImageID)
	if err != nil {
		return response.ResV12RecommendFood{}, err
	}
	imageUrl, err := aws.ImageGetSignedURL(ctx, image, aws.ImgTypeFood)
	if err != nil {
		return response.ResV12RecommendFood{}, err
	}
	nutrientDTO, err := d.Repository.FindOneNutrient(ctx, food.Name)
	if err != nil {
		return response.ResV12RecommendFood{}, err
	}
	res := CreateRes12Recommend(food, imageUrl, nutrientDTO)
	return res, nil
}
