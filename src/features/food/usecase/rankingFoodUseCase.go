package usecase

import (
	"context"

	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	_redis "main/utils/db/redis"
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
	currentRankings := make([]*entity.RankFoodRedis, 0)
	var err error
	currentRankings, err = d.Repository.RankingTop(ctx)
	if err != nil {
		return response.ResRankingFood{}, err
	}
	//현재 랭킹 데이터가 비어 있다면
	if len(currentRankings) == 0 {
		//rdb에서 데이터를 가져와 현재 랭킹에 저장을 한다.
		currentRankings, err = d.Repository.FindRankingFoodHistories(ctx)
		if err != nil {
			return response.ResRankingFood{}, err
		}
		//현재 랭킹 레디스에 저장한다.
		for _, food := range currentRankings {
			err := d.Repository.IncrementFoodRanking(ctx, _redis.RankingKey, food.Name, food.Score)
			if err != nil {
				return response.ResRankingFood{}, err
			}
		}
	}

	//이전 데이터가 있다면 랭킹 변동을 계산한다.
	res := response.ResRankingFood{}
	for i, food := range currentRankings {
		if i == 10 {
			break
		}
		rank := i + 1
		rankFood := response.RankFood{
			Rank:       rank,
			Name:       food.Name,
			RankChange: "new",
		}

		res.Foods = append(res.Foods, rankFood)
	}

	return res, nil
}
