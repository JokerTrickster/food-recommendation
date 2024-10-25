package repository

import (
	"context"
	"fmt"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewImageUploadFoodRepository(gormDB *gorm.DB) _interface.IImageUploadFoodRepository {
	return &ImageUploadFoodRepository{GormDB: gormDB}
}

func (g *ImageUploadFoodRepository) FindOneAndUpdateFoodImages(ctx context.Context, foodName, fileName string) error {
	fmt.Println(foodName)
	result := mysql.GormMysqlDB.WithContext(ctx).Model(&mysql.FoodImages{}).Where("name = ?", foodName).Update("image", fileName)
	// 오류 체크
	if result.Error != nil {
		fmt.Println(result.Error)
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(result.Error.Error(), foodName, fileName), utils.ErrFromMysqlDB)
	}

	// 업데이트된 행 수 확인
	if result.RowsAffected == 0 {
		return utils.ErrorMsg(ctx, utils.ErrFoodNotFound, utils.Trace(), "no matching record found", utils.ErrFromClient)
	}
	return nil
}
