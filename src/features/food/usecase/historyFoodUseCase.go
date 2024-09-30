package usecase

import (
	"context"

	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/features/food/model/response"
	"main/utils"
	"main/utils/db/mysql"
	"time"
)

type HistoryFoodUseCase struct {
	Repository     _interface.IHistoryFoodRepository
	ContextTimeout time.Duration
}

func NewHistoryFoodUseCase(repo _interface.IHistoryFoodRepository, timeout time.Duration) _interface.IHistoryFoodUseCase {
	return &HistoryFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *HistoryFoodUseCase) History(c context.Context, userID uint) (response.ResHistoryFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//db에서 조회한다.
	foodHistoryList, err := d.Repository.FindAllFoodHistory(ctx, userID)
	if err != nil {
		return response.ResHistoryFood{}, err
	}

	res := response.ResHistoryFood{}
	foods := make([]response.HistoryFood, 0)
	//TODO 추후 성능 처리 개선 필요
	for _, foodHistory := range foodHistoryList {
		foodDTO, err := d.Repository.FindOneFood(ctx, foodHistory.FoodID)
		if err != nil {
			return response.ResHistoryFood{}, err
		}
		timeID, ok := mysql.GetTimeKey(foodDTO.TimeID)
		if !ok {
			return response.ResHistoryFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+" timeID is not valid", utils.ErrFromInternal)
		}
		typeID, ok := mysql.GetTypeKey(foodDTO.TypeID)
		if !ok {
			return response.ResHistoryFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+" typeID is not valid", utils.ErrFromInternal)
		}
		scenarioID, ok := mysql.GetScenarioKey(foodDTO.ScenarioID)
		if !ok {
			return response.ResHistoryFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+" scenarioID is not valid", utils.ErrFromInternal)
		}
		themeID, ok := mysql.GetThemeKey(foodDTO.ThemeID)
		if !ok {
			return response.ResHistoryFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+" themeID is not valid", utils.ErrFromInternal)
		}
		flavorID, ok := mysql.GetFlavorKey(foodDTO.FlavorID)
		if !ok {
			return response.ResHistoryFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+" flavorID is not valid", utils.ErrFromInternal)
		}

		food := response.HistoryFood{
			Name:     foodDTO.Name,
			Type:     typeID,
			Time:     timeID,
			Scenario: scenarioID,
			Theme:    themeID,
			Flavor:   flavorID,
			Created:  foodHistory.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		foods = append(foods, food)
	}
	res.Foods = foods

	return res, nil
}
