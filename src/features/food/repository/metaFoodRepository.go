package repository

import (
	"context"
	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewMetaFoodRepository(gormDB *gorm.DB) _interface.IMetaFoodRepository {
	return &MetaFoodRepository{GormDB: gormDB}
}

func (g *MetaFoodRepository) FindAllTypeMeta(ctx context.Context) ([]mysql.Types, error) {
	var typeDTO []mysql.Types
	if err := g.GormDB.WithContext(ctx).Find(&typeDTO).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return typeDTO, nil
}
func (g *MetaFoodRepository) FindAllTimeMeta(ctx context.Context) ([]mysql.Times, error) {
	var timeDTO []mysql.Times
	if err := g.GormDB.WithContext(ctx).Find(&timeDTO).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return timeDTO, nil
}

func (g *MetaFoodRepository) FindAllScenarioMeta(ctx context.Context) ([]mysql.Scenarios, error) {
	var scenarioDTO []mysql.Scenarios
	if err := g.GormDB.WithContext(ctx).Find(&scenarioDTO).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return scenarioDTO, nil
}

func (g *MetaFoodRepository) FindAllThemesMeta(ctx context.Context) ([]mysql.Themes, error) {
	var themesDTO []mysql.Themes
	if err := g.GormDB.WithContext(ctx).Find(&themesDTO).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return themesDTO, nil
}

func (g *MetaFoodRepository) FindAllFlavorMeta(ctx context.Context) ([]mysql.Flavors, error) {
	var flavorDTO []mysql.Flavors
	if err := g.GormDB.WithContext(ctx).Find(&flavorDTO).Error; err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromMysqlDB)
	}
	return flavorDTO, nil
}
