package _interface

import (
	"context"
	"main/features/food/model/entity"
)

type IRecommendFoodUseCase interface {
	Recommend(c context.Context, entity entity.RecommendFoodEntity) ([]string,error)
}
