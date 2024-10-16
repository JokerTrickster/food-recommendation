package usecase

import (
	"context"

	_interface "main/features/food/model/interface"
	"main/features/food/model/request"
	"main/utils/aws"
	"time"
)

type SaveFoodUseCase struct {
	Repository     _interface.ISaveFoodRepository
	ContextTimeout time.Duration
}

func NewSaveFoodUseCase(repo _interface.ISaveFoodRepository, timeout time.Duration) _interface.ISaveFoodUseCase {
	return &SaveFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *SaveFoodUseCase) Save(c context.Context, req *request.ReqSaveFood) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	foods := make([]string, 0)
	for _, food := range req.Foods {
		foodImageDTO := CreateSaveFoodImageDTO(food)
		foodImage, err := d.Repository.FindOneOrCreateFoodImage(ctx, foodImageDTO)
		if err != nil {
			return err
		}
		foodDTO := CreateSaveFoodDTO(food, int(foodImage.ID))
		err = d.Repository.SaveFood(ctx, foodDTO)
		if err != nil {
			return err
		}
		foods = append(foods, food.Name)
	}

	go aws.EmailSendFoodNameReport(foods)
	return nil
}
