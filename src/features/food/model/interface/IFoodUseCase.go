package _interface

import (
	"context"
	"main/features/food/model/entity"
	"main/features/food/model/request"
	"main/features/food/model/response"
)

type IRecommendFoodUseCase interface {
	Recommend(c context.Context, entity entity.RecommendFoodEntity) (response.ResRecommendFood, error)
}
type ISelectFoodUseCase interface {
	Select(c context.Context, entity entity.SelectFoodEntity) (response.ResSelectFood, error)
}

type IHistoryFoodUseCase interface {
	History(c context.Context, userID uint) (response.ResHistoryFood, error)
}

type IMetaFoodUseCase interface {
	Meta(c context.Context) (response.ResMetaData, error)
}

type IRankingFoodUseCase interface {
	Ranking(c context.Context) (response.ResRankingFood, error)
}

type IImageUploadFoodUseCase interface {
	ImageUpload(c context.Context, e entity.ImageUploadFoodEntity) error
}
type IEmptyImageFoodUseCase interface {
	EmptyImage(c context.Context) (response.ResEmptyImageFood, error)
}

type IDailyRecommendFoodUseCase interface {
	DailyRecommend(c context.Context) (response.ResDailyRecommendFood, error)
}

type ISaveFoodUseCase interface {
	Save(c context.Context, req *request.ReqSaveFood) error
}
