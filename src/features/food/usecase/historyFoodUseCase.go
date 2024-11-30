package usecase

import (
	"context"

	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	"time"
)

type HistoryFoodUseCase struct {
	Repository     _interface.IHistoryFoodRepository
	ContextTimeout time.Duration
}

func NewHistoryFoodUseCase(repo _interface.IHistoryFoodRepository, timeout time.Duration) _interface.IHistoryFoodUseCase {
	return &HistoryFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *HistoryFoodUseCase) History(c context.Context, userID uint) (response.ResHistoryFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//db에서 조회한다.
	foodHistoryList, err := d.Repository.FindAllFoodHistory(ctx, userID)
	if err != nil {
		return response.ResHistoryFood{}, err
	}

	res := response.ResHistoryFood{}
	foods := make([]response.HistoryFood, 0)
	//TODO 추후 성능 처리 개선 필요
	for _, foodHistory := range foodHistoryList {
		foodDTO, err := d.Repository.FindOneFood(ctx, uint(foodHistory.FoodID))
		if err != nil {
			return response.ResHistoryFood{}, err
		}

		food := response.HistoryFood{
			Name:    foodDTO.Name,
			Created: foodHistory.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		foods = append(foods, food)
	}
	res.Foods = foods

	return res, nil
}
