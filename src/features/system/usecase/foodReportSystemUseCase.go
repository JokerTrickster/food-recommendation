package usecase

import (
	"context"
	_interface "main/features/system/model/interface"
	"main/utils/aws"
	"time"
)

type FoodReportSystemUseCase struct {
	Repository     _interface.IFoodReportSystemRepository
	ContextTimeout time.Duration
}

func NewFoodReportSystemUseCase(repo _interface.IFoodReportSystemRepository, timeout time.Duration) _interface.IFoodReportSystemUseCase {
	return &FoodReportSystemUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *FoodReportSystemUseCase) FoodReport(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//음식 이미지가 없는 음식 리스트 조회
	images, err := d.Repository.FindFoodWithoutImage(ctx)
	if err != nil {
		return err
	}

	//음식 영양성분이 없는 음식 리스트 조회
	nutrients, err := d.Repository.FindFoodWithoutNutrient(ctx)
	if err != nil {
		return err
	}
	imageList := make([]string, 0)
	nutrientList := make([]string, 0)
	for _, image := range images {
		imageList = append(imageList, image.Name)
	}
	for _, nutrient := range nutrients {
		nutrientList = append(nutrientList, nutrient.Name)
	}

	go aws.EmailSendFoodInfoEmptyReport(imageList, nutrientList)

	return nil
}
