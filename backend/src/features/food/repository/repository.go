package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RecommendFoodRepository struct {
	GormDB *gorm.DB
}

type SelectFoodRepository struct {
	GormDB      *gorm.DB
	RedisClient *redis.Client
}

type HistoryFoodRepository struct {
	GormDB *gorm.DB
}
type MetaFoodRepository struct {
	GormDB *gorm.DB
}

type RankingFoodRepository struct {
	GormDB      *gorm.DB
	RedisClient *redis.Client
}
