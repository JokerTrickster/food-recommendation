package _interface

import (
	"context"
	"main/features/food/model/entity"
	"main/utils/db/mysql"
)

type IRecommendFoodRepository interface {
	SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) (*mysql.Foods, error)
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
	FindAllRanking(ctx context.Context, redisKey string) ([]string, error)
	FindPreviousRanking(ctx context.Context, todayRedisKey, yesterDayRedisKey string, food string, currentRank int) (string, error)
	RankingTop(ctx context.Context) ([]*entity.RankFoodRedis, error)
	PreviousRanking(ctx context.Context, food string) (int, error)
	FindRankingTop10FoodHistories(ctx context.Context) ([]*entity.RankFoodRedis, error)
	IncrementFoodRanking(ctx context.Context, redisKey string, foodName string, score float64) error
	PreviousRankingExist(ctx context.Context) (int, error)
	FindOneFoods(ctx context.Context, foodID int) (string, error)
	ExpireRanking(ctx context.Context, key string) error
}

type IImageUploadFoodRepository interface {
	FindOneAndUpdateFoods(ctx context.Context, foodID uint, filename string) error
}

type IEmptyImageFoodRepository interface {
	FindAllEmptyImageFoods(ctx context.Context) ([]mysql.Foods, error)
}

type IDailyRecommendFoodRepository interface {
	FindOneFood(ctx context.Context, foodName string) (*mysql.Foods, error)
}
