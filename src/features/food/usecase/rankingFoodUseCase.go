package usecase

import (
	"context"
	"strconv"

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
		currentRankings, err = d.Repository.FindRankingTop10FoodHistories(ctx)
		if err != nil {
			return response.ResRankingFood{}, err
		}
		//현재 랭킹 레디스에 저장한다.
		for _, food := range currentRankings {
			err := d.Repository.IncrementFoodRanking(ctx, _redis.RankingKey, food.FoodID, 1)
			if err != nil {
				return response.ResRankingFood{}, err
			}
		}
	}

	// 이전 데이터가 비어 있다면 현재 랭킹 데이터를 저장한다.
	previousExist, err := d.Repository.PreviousRankingExist(ctx)
	if err != nil {
		return response.ResRankingFood{}, err
	}
	if previousExist == 0 {
		//이전 데이터가 없다면 현재 랭킹 데이터를 이전 랭킹 데이터로 저장한다.
		for _, food := range currentRankings {
			err := d.Repository.IncrementFoodRanking(ctx, _redis.PrevRankingKey, food.FoodID, food.Score)
			if err != nil {
				return response.ResRankingFood{}, err
			}
		}
		//이전 데이터 만료시간을 설정한다.
		err := d.Repository.ExpireRanking(ctx, _redis.PrevRankingKey)
		if err != nil {
			return response.ResRankingFood{}, err
		}
	}
	//이전 데이터가 있다면 랭킹 변동을 계산한다.
	res := response.ResRankingFood{}
	for i, food := range currentRankings {
		if i == 10 {
			break
		}
		rank := i + 1
		previousRank, err := d.Repository.PreviousRanking(ctx, food.FoodID)
		if err != nil {
			return response.ResRankingFood{}, err
		}
		fID, _ := strconv.Atoi(food.FoodID)
		foodName, err := d.Repository.FindOneFoods(ctx, fID)
		if err != nil {
			return response.ResRankingFood{}, err
		}
		rankFood := response.RankFood{
			Rank: rank,
			Name: foodName,
		}
		if previousRank == _redis.NewRank {
			rankFood.RankChange = "new"
		} else {
			rankChange := previousRank - rank
			rankFood.RankChange = strconv.Itoa(rankChange)
		}

		res.Foods = append(res.Foods, rankFood)
	}

	return res, nil
}
