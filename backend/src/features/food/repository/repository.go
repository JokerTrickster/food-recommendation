package repository

import "gorm.io/gorm"

type RecommendFoodRepository struct {
	GormDB *gorm.DB
}
