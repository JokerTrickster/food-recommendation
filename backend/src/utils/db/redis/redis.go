package _redis

import (
	"context"
	"fmt"
	"os"

	_aws "main/utils/aws"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

const RankingKey = "food:rankings"

func InitRedis() error {
	ctx := context.Background()
	isLocal := os.Getenv("IS_LOCAL")
	var connectionString string
	if isLocal == "true" {
		connectionString = fmt.Sprintf("redis://%s:%s@localhost:6379/0", os.Getenv("REDIS_USER"), os.Getenv("REDIS_PASSWORD"))
	} else {
		dbInfos, err := _aws.AwsSsmGetParams([]string{"dev_food-recommendation_redis_user", "dev_food-recommendation_redis_password", "dev_food-recommendation_redis_host", "dev_food-recommendation_redis_port", "dev_food-recommendation_redis_db"})
		if err != nil {
			return err
		}
		connectionString = fmt.Sprintf("redis://%s:%s@%s:%s/%s",
			dbInfos[4], //user
			dbInfos[1], //password
			dbInfos[3], //host
			dbInfos[2], //port
			dbInfos[0], //db name
		)
	}

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		return err
	}

	Client = redis.NewClient(opt)

	_, err = Client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("Connected to Redis!")

	return nil
}
