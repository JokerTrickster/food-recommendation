package usecase

import (
	"context"
	"fmt"

	_interface "main/features/food/model/interface"
	"main/features/food/model/response"

	"time"
)

type EmptyImageFoodUseCase struct {
	Repository     _interface.IEmptyImageFoodRepository
	ContextTimeout time.Duration
}

func NewEmptyImageFoodUseCase(repo _interface.IEmptyImageFoodRepository, timeout time.Duration) _interface.IEmptyImageFoodUseCase {
	return &EmptyImageFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *EmptyImageFoodUseCase) EmptyImage(c context.Context) (response.ResEmptyImageFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Println(ctx)

	foods, err := d.Repository.FindAllEmptyImageFoods(ctx)
	if err != nil {
		return response.ResEmptyImageFood{}, err
	}
	res := CreateResEmptyImageFood(foods)
	return res, nil
}
