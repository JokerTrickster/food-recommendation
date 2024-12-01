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
	InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *mysql.FoodHistories) error
	IncrementFoodRanking(ctx context.Context, foodName string, score float64) error
}
type IHistoryFoodRepository interface {
	FindAllFoodHistory(ctx context.Context, userID uint) ([]mysql.FoodHistories, error)
	FindOneFood(ctx context.Context, foodID uint) (*mysql.Foods, error)
}

type IMetaFoodRepository interface {
	FindAllTypeMeta(ctx context.Context) ([]mysql.Types, error)
	FindAllTimeMeta(ctx context.Context) ([]mysql.Times, error)
	FindAllScenarioMeta(ctx context.Context) ([]mysql.Scenarios, error)
	FindAllThemesMeta(ctx context.Context) ([]mysql.Themes, error)
}

type IRankingFoodRepository interface {
	RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error)
	FindRankingFoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error)
	IncrementFoodRanking(ctx context.Context, redisKey string, foodName string, score float64) error
}

type IImageUploadFoodRepository interface {
	FindOneAndUpdateFoodImages(ctx context.Context, foodName, fileName string) error
}

type IEmptyImageFoodRepository interface {
	FindAllEmptyImageFoods(ctx context.Context) ([]mysql.FoodImages, error)
}

type IDailyRecommendFoodRepository interface {
	FindOneFood(ctx context.Context, foodName string) (*mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindRandomFoods(ctx context.Context, limit int) ([]*mysql.Foods, error)
}

type ISaveFoodRepository interface {
	SaveFood(ctx context.Context, foodDTO *mysql.Foods) (uint, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *mysql.FoodImages) (*mysql.FoodImages, error)
	SaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) error
	FindCategoryIDs(ctx context.Context, categories []string) ([]uint, error)
	SaveFoodCategory(ctx context.Context, foodID uint, categoryIDs []uint) error
}

type ICheckImageUploadFoodRepository interface {
}

type IV1RecommendFoodRepository interface {
	FindOneV1RecommendFood(ctx context.Context, query string) (*mysql.Foods, error)
	SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) (*mysql.Foods, error)
	FindOneOrCreateFoodImage(ctx context.Context, foodImageDTO *mysql.FoodImages) (*mysql.FoodImages, error)
	CountV1RecommendFood(ctx context.Context, query string) (int, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindOneNutrient(ctx context.Context, foodName string) (*mysql.Nutrients, error)
	FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) (*mysql.Nutrients, error)
}

type IV12RecommendFoodRepository interface {
	FindOneV12RecommendFood(ctx context.Context, query string) (*mysql.Foods, error)
	FindOneFoodImage(ctx context.Context, foodID int) (string, error)
	FindOneNutrient(ctx context.Context, foodName string) (*mysql.Nutrients, error)
	FindOneAndSaveNutrient(ctx context.Context, nutrientDTO *mysql.Nutrients) (*mysql.Nutrients, error)
}
