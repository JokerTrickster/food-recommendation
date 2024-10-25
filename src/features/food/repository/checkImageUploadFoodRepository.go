package repository

import (
	_interface "main/features/food/model/interface"

	"gorm.io/gorm"
)

func NewCheckImageUploadFoodRepository(gormDB *gorm.DB) _interface.ICheckImageUploadFoodRepository {
	return &CheckImageUploadFoodRepository{GormDB: gormDB}
}
