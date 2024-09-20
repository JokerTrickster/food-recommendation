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
	res := response.ResRankingFood{}
	// 순위 업데이트 및 저장
	for i, food := range foodRank {
		rankFood := response.RankFood{}
		currentRank := i + 1 // 1위부터 시작

		// 이전 순위 가져오기
		rankChange, err := d.Repository.FindPreviousRanking(ctx, food, int(currentRank))
		if err != nil {
			continue
		}
		rankFood.Name = food
		rankFood.Rank = int(currentRank)
		rankFood.RankChange = rankChange
		res.Foods = append(res.Foods, rankFood)
	}

	return res, nil
}
