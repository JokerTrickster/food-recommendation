package usecase

import (
	"context"
	"fmt"

	_interface "main/features/food/model/interface"
	"time"
)

type SaveFoodUseCase struct {
	Repository     _interface.ISaveFoodRepository
	ContextTimeout time.Duration
}

func NewSaveFoodUseCase(repo _interface.ISaveFoodRepository, timeout time.Duration) _interface.ISaveFoodUseCase {
	return &SaveFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SaveFoodUseCase) Save(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	fmt.Println(ctx)
	return nil
}
