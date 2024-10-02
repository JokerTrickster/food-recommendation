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
		genai.Text("한국에 살고 있는 사람들에게 맛있는 음식을 추천해주는 전문가이다."),
		genai.Text("상황을 제시해주면 상황에 맞는 맛있는 음식을 추천해주면 된다."),
		genai.Text("반드시 음식 이름 1개만 추천해줘야 되며, 요리법, 재료, 가게 이름 등으로 대답해주면 안된다."),
		genai.Text("예를들면 오늘 친구와 같이 점심에 매운 음식 이름을 추천받고 싶다면 대답으로 닭갈비 라고 하면 된다."),
		genai.Text("응답을 해줄때 음식 이름인지 한번 더 확인 후 대답해줘"),
		genai.Text("음식 이름을 응답해줄 때 이모티콘,특수문자(*&^$@~!@...) 등을 포함해서 대답해주면 안된다."),
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
		if err != nil {
			return response.ResRecommendFood{}, err
		}
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
