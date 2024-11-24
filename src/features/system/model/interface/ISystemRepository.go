package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IReportSystemRepository interface {
	SaveReport(ctx context.Context, reportDTO *mysql.Reports) error
}

type IFoodReportSystemRepository interface {
	FindFoodWithoutImage(ctx context.Context) ([]*mysql.FoodImages, error)
	FindFoodWithoutNutrient(ctx context.Context) ([]*mysql.Foods, error)
}
