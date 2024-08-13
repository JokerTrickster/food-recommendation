package usecase

import (
	"context"

	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	"time"
)

type RankingFoodUseCase struct {
	Repository     _interface.IRankingFoodRepository
	ContextTimeout time.Duration
}

func NewRankingFoodUseCase(repo _interface.IRankingFoodRepository, timeout time.Duration) _interface.IRankingFoodUseCase {
	return &RankingFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RankingFoodUseCase) Ranking(c context.Context) (response.ResRankingFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	foodRank, err := d.Repository.FindAllRanking(ctx)
	if err != nil {
		return response.ResRankingFood{}, err
	}

	res := CreateResRankingFood(foodRank)

	return res, nil
}
