package repository

import (
	"context"
	"fmt"
	_interface "main/features/food/model/interface"
	_redis "main/utils/db/redis"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRankingFoodRepository(gormDB *gorm.DB, redisClient *redis.Client) _interface.IRankingFoodRepository {
	return &RankingFoodRepository{GormDB: gormDB, RedisClient: redisClient}
}

func (g *RankingFoodRepository) FindAllRanking(ctx context.Context) ([]string, error) {
	today := time.Now().Format("2006-01-02")
	redisKey := _redis.RankingKey + ":" + today
	foodList, err := g.RedisClient.ZRevRange(ctx, redisKey, 0, 9).Result()
	if err != nil {
		return nil, err
	}
	return foodList, nil
}

func (g *RankingFoodRepository) FindPreviousRanking(ctx context.Context, food string, currentRank int) (string, error) {
	yesterDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	redisKey := _redis.RankingKey + ":" + yesterDay
	prevRank, err := g.RedisClient.HGet(ctx, redisKey, food).Result()
	if err == redis.Nil {
		return "new", nil
	} else if err != nil {
		fmt.Println("Error fetching previous ranking:", err)
		return "", err
	} else {
		fmt.Println(prevRank)
		// 이전 순위가 존재하는 경우 변동 계산
		intPrevRank, err := strconv.Atoi(prevRank)
		if err != nil {
			return "", err
		}
		rankChange := intPrevRank - currentRank
		rankChangeStr := strconv.Itoa(rankChange)
		return rankChangeStr, nil
	}
}

func (g *RankingFoodRepository) SavePreviousRanking(ctx context.Context, food string, currentRank int) error {
	yesterDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	redisKey := _redis.RankingKey + ":" + yesterDay
	err := g.RedisClient.HSet(ctx, redisKey, food, currentRank).Err()
	if err != nil {
		return err
	}
	return nil
}
