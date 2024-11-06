package usecase

import (
	"context"
	"main/features/food/model/entity"
	"main/features/food/model/response"
	"main/utils/aws"

	_interface "main/features/food/model/interface"
	"time"
)

type V1RecommendFoodUseCase struct {
	Repository     _interface.IV1RecommendFoodRepository
	ContextTimeout time.Duration
}

func NewV1RecommendFoodUseCase(repo _interface.IV1RecommendFoodRepository, timeout time.Duration) _interface.IV1RecommendFoodUseCase {
	return &V1RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}
func (d *V1RecommendFoodUseCase) V1Recommend(c context.Context, e entity.V1RecommendFoodEntity) (response.ResV1RecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// 람다 api 호출
	lambdaURL, err := aws.AwsSsmGetParam("food_recommend_name_api")
	if err != nil {
		return response.ResV1RecommendFood{}, err
	}
	requestBody := []byte(`{"key": "value"}`)

	err = HttpCallRecommendNameApi(ctx, lambdaURL, requestBody)
	if err != nil {
		return response.ResV1RecommendFood{}, err
	}

	return response.ResV1RecommendFood{}, nil
}
