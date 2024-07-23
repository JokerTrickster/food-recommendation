package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IRecommendFoodRepository interface {
	SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) error
}

type ISelectFoodRepository interface {
	FindOneFood(ctx context.Context, foodDTO *mysql.Foods) (uint, error)
	InsertOneFoodHistory(ctx context.Context, foodHistoryDTO *mysql.FoodHistory) error
}
type IHistoryFoodRepository interface {
	FindAllFoodHistory(ctx context.Context, userID uint) ([]mysql.FoodHistory, error)
	FindOneFood(ctx context.Context, foodID uint) (*mysql.Foods, error)
}

type IMetaFoodRepository interface {
	FindAllTypeMeta(ctx context.Context) ([]mysql.Types, error)
	FindAllTimeMeta(ctx context.Context) ([]mysql.Times, error)
	FindAllScenarioMeta(ctx context.Context) ([]mysql.Scenarios, error)
}
