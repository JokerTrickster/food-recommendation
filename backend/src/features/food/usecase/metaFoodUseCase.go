package usecase

import (
	"context"
	"fmt"

	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	"time"
)

type MetaFoodUseCase struct {
	Repository     _interface.IMetaFoodRepository
	ContextTimeout time.Duration
}

func NewMetaFoodUseCase(repo _interface.IMetaFoodRepository, timeout time.Duration) _interface.IMetaFoodUseCase {
	return &MetaFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *MetaFoodUseCase) Meta(c context.Context) (response.ResMetaData, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Println(ctx)
	typeDTO, err := d.Repository.FindAllTypeMeta(ctx)
	if err != nil {
		return response.ResMetaData{}, err
	}
	timeDTO, err := d.Repository.FindAllTimeMeta(ctx)
	if err != nil {
		return response.ResMetaData{}, err
	}
	scenarioDTO, err := d.Repository.FindAllScenarioMeta(ctx)
	if err != nil {
		return response.ResMetaData{}, err
	}
	res := CreateResMetaData(typeDTO, timeDTO, scenarioDTO)

	return res, nil
}
