package _interface

import (
	"context"
	"main/features/food/model/entity"
	"main/features/food/model/response"
)

type IRecommendFoodUseCase interface {
	Recommend(c context.Context, entity entity.RecommendFoodEntity) ([]string, error)
}
type ISelectFoodUseCase interface {
	Select(c context.Context, entity entity.SelectFoodEntity) error
}

type IHistoryFoodUseCase interface {
	History(c context.Context, userID uint) (response.ResHistoryFood, error)
}

type IMetaFoodUseCase interface {
	Meta(c context.Context) (response.ResMetaData, error)
}
