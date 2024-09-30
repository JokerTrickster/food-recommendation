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

type RecommendFoodUseCase struct {
	Repository     _interface.IRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewRecommendFoodUseCase(repo _interface.IRecommendFoodRepository, timeout time.Duration) _interface.IRecommendFoodUseCase {
	return &RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RecommendFoodUseCase) Recommend(c context.Context, e entity.RecommendFoodEntity) (response.ResRecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//음식 추천 로직 구현
	client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GeminiID))
	if err != nil {
		return response.ResRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	//데이터 가공
	question := CreateRecommendFoodQuestion(e)
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("너는 맛있는 요리 음식 이름을 알려주는 전문가이다."),
		genai.Text("내가 질문을 하면 음식 이름으로만 대답을 해줘야 된다."),
		genai.Text("예를 들어서 '매운 음식 추천해줘' 라고 물으면 '김치찌개' 라고 대답을 해줘야 된다."),
		genai.Text("반드시 음식 이름 1개만 추천해줘, 요리법, 재료 등 다른 정보는 필요없다"),
		genai.Text("응답을 해줄때 음식 이름인지 한번더 생각하고 말해줘"),
		genai.Text("지금부터 질문할게 대답해줘"),
		genai.Text(question),
	)

	if err != nil {
		return response.ResRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	gptRes := make([]string, 0)
	// 출력 부분 수정

	if len(resp.Candidates) > 0 {
		marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
		var generateResponse entity.ContentResponse
		if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
			return response.ResRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromInternal)
		}
		for _, cad := range *generateResponse.Candidates {
			if cad.Content != nil {
				cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
				gptRes = SplitAndRemoveEmpty(cleanedString)
			}
		}

	} else {
		return response.ResRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrNotFound, utils.Trace(), _errors.ErrFoodNotFound.Error(), utils.ErrFromGemini)
	}
	res := response.ResRecommendFood{}
	//db에 저장
	for _, foodName := range gptRes {
		foodImageDTO := CreateRecommendFoodImageDTO(e, foodName)
		foodImage, err := d.Repository.FindOneOrCreateFoodImage(ctx, foodImageDTO)
		foodDTO := CreateRecommendFoodDTO(e, foodName, int(foodImage.ID))

		foods, err := d.Repository.SaveRecommendFood(ctx, foodDTO)
		if err != nil {
			return response.ResRecommendFood{}, err
		}
		food := response.RecommendFood{
			Name: foods.Name,
		}
		imageUrl, err := aws.ImageGetSignedURL(ctx, foodImage.Image, aws.ImgTypeFood)
		if err != nil {
			return response.ResRecommendFood{}, err
		}
		food.Image = imageUrl
		res.FoodNames = append(res.FoodNames, food)
		break
	}
	return res, nil
}
