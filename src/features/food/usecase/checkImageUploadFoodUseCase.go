package usecase

import (
	"context"

	_interface "main/features/food/model/interface"
	"main/features/food/model/request"
	"main/utils/aws"

	"time"
)

type CheckImageUploadFoodUseCase struct {
	Repository     _interface.ICheckImageUploadFoodRepository
	ContextTimeout time.Duration
}

func NewCheckImageUploadFoodUseCase(repo _interface.ICheckImageUploadFoodRepository, timeout time.Duration) _interface.ICheckImageUploadFoodUseCase {
	return &CheckImageUploadFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *CheckImageUploadFoodUseCase) CheckImageUpload(c context.Context, req *request.ReqCheckImageUploadFood) error {
	_, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	aws.EmailSendFoodUploadReport(req.SuccessFoodList, req.FailedFoodList)

	return nil
}
