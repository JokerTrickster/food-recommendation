package usecase

import (
	"context"

	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"time"
)

type SelectFoodUseCase struct {
	Repository     _interface.ISelectFoodRepository
	ContextTimeout time.Duration
}

func NewSelectFoodUseCase(repo _interface.ISelectFoodRepository, timeout time.Duration) _interface.ISelectFoodUseCase {
	return &SelectFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SelectFoodUseCase) Select(c context.Context, e entity.SelectFoodEntity) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//db에서 조회한다.
	foodDTO := CreateSelectFoodDTO(e)
	foodID, err := d.Repository.FindOneFood(ctx, foodDTO)
	if err != nil {
		return err
	}
	foodDTO.ID = foodID

	//디비에 저장한다.
	foodHistoryDTO := CreateFoodHistoryDTO(foodID, e.UserID)
	if err := d.Repository.InsertOneFoodHistory(ctx, foodHistoryDTO); err != nil {
		return err
	}

	//레디스 저장한다.
	if err := d.Repository.IncrementFoodRanking(ctx, foodDTO.Name, 1); err != nil {
		return err
	}

	return nil
}
