package usecase

import (
	"context"
	"fmt"

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
		//음식 이미지 저장
		foodImageDTO := CreateSaveFoodImageDTO(food)
		foodImage, err := d.Repository.FindOneOrCreateFoodImage(ctx, foodImageDTO)
		if err != nil {
			return err
		}
		//음식 저장
		foodDTO := CreateSaveFoodDTO(food, int(foodImage.ID))
		foodID, err := d.Repository.SaveFood(ctx, foodDTO)
		if err != nil {
			return err
		}
		//카테고리 ID를 모두 가져온다.
		categories := CreateCategory(food)
		categoryIDs, err := d.Repository.FindCategoryIDs(ctx, categories)
		if err != nil {
			return err
		}
		//음식 카테고리 저장
		err = d.Repository.SaveFoodCategory(ctx, foodID, categoryIDs)
		if err != nil {
			return err
		}

		//영양성분 저장
		if food.Amount != "" {
			nutirentDTO := CreateSaveNutrientDTO(food)
			err = d.Repository.SaveNutrient(ctx, nutirentDTO)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		// 카테고리 저장
		foods = append(foods, food.Name)
	}

	go aws.EmailSendFoodNameReport(foods)
	return nil
}
