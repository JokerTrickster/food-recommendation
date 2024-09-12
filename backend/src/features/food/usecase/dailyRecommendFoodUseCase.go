package usecase

import (
	"context"
	"encoding/json"
	"main/features/food/model/entity"
	"main/features/food/model/response"
	"main/utils/aws"

	_errors "main/features/food/model/errors"
	_interface "main/features/food/model/interface"
	"main/utils"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type DailyRecommendFoodUseCase struct {
	Repository     _interface.IDailyRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewDailyRecommendFoodUseCase(repo _interface.IDailyRecommendFoodRepository, timeout time.Duration) _interface.IDailyRecommendFoodUseCase {
	return &DailyRecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *DailyRecommendFoodUseCase) DailyRecommend(c context.Context) (response.ResDailyRecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//음식 추천 로직 구현
	client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GeminiID))
	if err != nil {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	//데이터 가공
	question := CreateDailyRecommendFoodQuestion()
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("너는 맛있는 요리 음식을 알려주는 전문가이다."),
		genai.Text("내가 질문을 하면 단어로만 대답을 해줘야 된다."),
		genai.Text("예를 들어서 '매운 음식 추천해줘' 라고 물으면 '김치찌개' 라고 대답을 해줘야 된다."),
		genai.Text("반드시 음식 이름 3개 추천해줘"),
		genai.Text("지금부터 질문할게 대답해줘 대답 해줄때 공백으로 구분지어서 대답해줘"),
		genai.Text(question),
	)

	if err != nil {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	gptRes := make([]string, 0)
	// 출력 부분 수정
	if len(resp.Candidates) > 0 {
		marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
		var generateResponse entity.ContentResponse
		if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
			return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromInternal)
		}
		for _, cad := range *generateResponse.Candidates {
			if cad.Content != nil {
				cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
				gptRes = SplitAndRemoveEmpty(cleanedString)
			}
		}

	} else {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrNotFound, utils.Trace(), _errors.ErrFoodNotFound.Error(), utils.ErrFromGemini)
	}
	res := response.ResDailyRecommendFood{}
	//db에서 가져온다.
	for _, foodName := range gptRes {
		food := response.DailyRecommendFood{
			Name:  foodName,
			Image: "food-default.png",
		}
		foods, err := d.Repository.FindOneFood(ctx, foodName)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		if foods != nil {
			food.Image = foods.Image
		}

		imageUrl, err := aws.ImageGetSignedURL(ctx, food.Image, aws.ImgTypeFood)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		food.Image = imageUrl
		res.DilayFoods = append(res.DilayFoods, food)
	}
	return res, nil
}
