package repository

import "gorm.io/gorm"

type RecommendFoodRepository struct {
	GormDB *gorm.DB
}

type SelectFoodRepository struct {
	GormDB *gorm.DB
}

type HistoryFoodRepository struct {
	GormDB *gorm.DB
}
type MetaFoodRepository struct {
	GormDB *gorm.DB
}
