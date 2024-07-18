package usecase

import (
	"context"
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"time"
)

type RecommendFoodUseCase struct {
	Repository     _interface.IRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewRecommendFoodUseCase(repo _interface.IRecommendFoodRepository, timeout time.Duration) _interface.IRecommendFoodUseCase {
	return &RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RecommendFoodUseCase) Recommend(c context.Context, entity entity.RecommendFoodEntity) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 유저 정보를 가져온다.

	//데이터 가공

	//음식 추천 로직 구현

	//db에 저장

	return nil
}
