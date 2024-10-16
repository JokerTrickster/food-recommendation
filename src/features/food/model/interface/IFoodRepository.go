package _interface

import (
	"context"
	"main/features/food/model/entity"
	"main/utils/db/mysql"
)

type IRecommendFoodRepository interface {
	SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) (*mysql.Foods, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *mysql.FoodImages) (*mysql.FoodImages, error)
}

type ISelectFoodRepository interface {
	FindOneFood(ctx context.Context, foodDTO *mysql.Foods) (uint, error)
	InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *mysql.FoodHistory) error
	IncrementFoodRanking(ctx context.Context, foodName string, score float64) error
}
type IHistoryFoodRepository interface {
	FindAllFoodHistory(ctx context.Context, userID uint) ([]mysql.FoodHistory, error)
	FindOneFood(ctx context.Context, foodID uint) (*mysql.Foods, error)
}

type IMetaFoodRepository interface {
	FindAllTypeMeta(ctx context.Context) ([]mysql.Types, error)
	FindAllTimeMeta(ctx context.Context) ([]mysql.Times, error)
	FindAllScenarioMeta(ctx context.Context) ([]mysql.Scenarios, error)
	FindAllThemesMeta(ctx context.Context) ([]mysql.Themes, error)
	FindAllFlavorMeta(ctx context.Context) ([]mysql.Flavors, error)
}

type IRankingFoodRepository interface {
	RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error)
	FindRankingFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error)
	IncrementFoodRanking(ctx context.Context, redisKey string, foodName string, score float64) error
}

type IImageUploadFoodRepository interface {
	FindOneAndUpdateFoodImages(ctx context.Context, foodID uint, filename string) error
}

type IEmptyImageFoodRepository interface {
	FindAllEmptyImageFoods(ctx context.Context) ([]mysql.FoodImages, error)
}

type IDailyRecommendFoodRepository interface {
	FindOneFood(ctx context.Context, foodName string) (*mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
}

type ISaveFoodRepository interface {
	SaveFood(ctx context.Context, foodDTO *[]mysql.Foods) error
}
