package _interface

import (
	"context"
	"main/utils/db/mysql"
)

type IRecommendFoodRepository interface {
	SaveRecommendFood(ctx context.Context, foodDTO *mysql.Foods) error
}
