package usecase

import (
	"context"
	"fmt"

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
	todayRedisKey := _redis.RankingKey + ":" + time.Now().Format("2006-01-02")
	foodRank, err := d.Repository.FindAllRanking(ctx, todayRedisKey)
	if err != nil {
		return response.ResRankingFood{}, err
	}
	res := response.ResRankingFood{}
	if len(foodRank) >= 10 {

		// 순위 업데이트 및 저장
		for i, food := range foodRank {
			rankFood := response.RankFood{}
			currentRank := i + 1 // 1위부터 시작
			yesterDayRedisKey := _redis.RankingKey + ":" + time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			// 이전 순위 가져오기
			rankChange, err := d.Repository.FindPreviousRanking(ctx, todayRedisKey, yesterDayRedisKey, food, int(currentRank))
			if err != nil {
				continue
			}
			rankFood.Name = food
			rankFood.Rank = int(currentRank)
			rankFood.RankChange = rankChange
			res.Foods = append(res.Foods, rankFood)
		}
	} else {
		// 순위 업데이트 및 저장
		prevDayRedisKey := _redis.RankingKey + ":" + time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		foodRank, err := d.Repository.FindAllRanking(ctx, prevDayRedisKey)
		if err != nil {
			return response.ResRankingFood{}, err
		}
		fmt.Println("foodRank", foodRank)
		for i, food := range foodRank {
			rankFood := response.RankFood{}
			currentRank := i + 1 // 1위부터 시작
			todayDayRedisKey := _redis.RankingKey + ":" + time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			yesterDayRedisKey := _redis.RankingKey + ":" + time.Now().AddDate(0, 0, -2).Format("2006-01-02")

			// 이전 순위 가져오기
			rankChange, err := d.Repository.FindPreviousRanking(ctx, todayDayRedisKey, yesterDayRedisKey, food, int(currentRank))
			if err != nil {
				continue
			}
			rankFood.Name = food
			rankFood.Rank = int(currentRank)
			rankFood.RankChange = rankChange
			res.Foods = append(res.Foods, rankFood)
		}
	}

	return res, nil
}
