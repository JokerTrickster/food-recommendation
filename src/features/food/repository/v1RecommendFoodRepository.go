package repository

import (
	_interface "main/features/food/model/interface"

	"gorm.io/gorm"
)

func NewV1RecommendFoodRepository(gormDB *gorm.DB) _interface.IV1RecommendFoodRepository {
	return &V1RecommendFoodRepository{GormDB: gormDB}
}
