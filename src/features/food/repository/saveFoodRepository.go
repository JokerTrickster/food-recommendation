package repository

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewSaveFoodRepository(gormDB *gorm.DB) _interface.ISaveFoodRepository {
	return &SaveFoodRepository{GormDB: gormDB}
}

func (d *SaveFoodRepository) SaveFood(ctx context.Context, foodDTO *[]mysql.Foods) error {
	return nil
}
