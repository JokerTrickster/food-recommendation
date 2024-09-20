package repository

import (
	"context"
	"fmt"
	_interface "main/features/food/model/interface"
	"main/utils"
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

	// 내림차순으로 10개 멤버와 점수 가져오기
	foodList, err := g.RedisClient.ZRevRangeWithScores(ctx, redisKey, 0, 9).Result()
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 음식 이름과 점수를 저장할 맵 생성
	rankings := make([]string, 0)

	// 결과를 맵에 추가
	for _, z := range foodList {
		rankings = append(rankings, z.Member.(string))
	}

	return rankings, nil
}
func (g *RankingFoodRepository) FindPreviousRanking(ctx context.Context, food string, currentRank int) (string, error) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")

	redisKeyYesterday := _redis.RankingKey + ":" + yesterday
	redisKeyToday := _redis.RankingKey + ":" + today

	// 어제의 점수를 가져오기
	_, err := g.RedisClient.ZScore(ctx, redisKeyYesterday, food).Result()
	if err == redis.Nil {
		return "new", nil // 이전 랭킹이 없으면 "new" 반환
	} else if err != nil {
		fmt.Println("Error fetching previous ranking:", err)
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 어제의 랭킹 가져오기
	prevRank, err := g.RedisClient.ZRevRank(ctx, redisKeyYesterday, food).Result()
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 오늘의 랭킹 가져오기
	currentRankToday, err := g.RedisClient.ZRevRank(ctx, redisKeyToday, food).Result()
	if err != nil {
		return "", utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), err.Error(), utils.ErrFromRedis)
	}

	// 랭킹 변동 계산
	rankChange := int(prevRank) - int(currentRankToday)
	rankChangeStr := strconv.Itoa(rankChange)

	return rankChangeStr, nil
}
