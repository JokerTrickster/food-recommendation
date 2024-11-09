package usecase

import (
	"context"
	"encoding/json"
	"main/features/food/model/entity"
	_errors "main/features/food/model/errors"
	"main/features/food/model/response"
	"main/utils"
	"main/utils/aws"
	"strings"

	_interface "main/features/food/model/interface"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type V1RecommendFoodUseCase struct {
	Repository     _interface.IV1RecommendFoodRepository
	ContextTimeout time.Duration
}

func NewV1RecommendFoodUseCase(repo _interface.IV1RecommendFoodRepository, timeout time.Duration) _interface.IV1RecommendFoodUseCase {
	return &V1RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}
func (d *V1RecommendFoodUseCase) V1Recommend(c context.Context, e entity.V1RecommendFoodEntity) (response.ResV1RecommendFood, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	query := CreateRecommendQuery(e)
	count, err := d.Repository.CountV1RecommendFood(ctx, query)
	if err != nil {
		return response.ResV1RecommendFood{}, err
	}

	// 음식 종류 수가 5가지 이하라면 AI로 추천
	if count < 5 {
		//음식 추천 로직 구현
		client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GeminiID))
		if err != nil {
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
		}
		model := client.GenerativeModel("gemini-1.5-flash")
		//데이터 가공
		question := CreateV1RecommendFoodQuestion(e)
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
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
		}
		gptRes := make([]string, 0)
		// 출력 부분 수정

		if len(resp.Candidates) > 0 {
			marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
			var generateResponse entity.ContentResponse
			if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
				return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromInternal)
			}
			for _, cad := range *generateResponse.Candidates {
				if cad.Content != nil {
					cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
					gptRes = SplitAndRemoveEmpty(cleanedString)
				}
			}

		} else {
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrGeminiError, utils.Trace(), _errors.ErrFoodNotFound.Error(), utils.ErrFromGemini)
		}
		aiRes := response.ResV1RecommendFood{}
		//db에 저장
		for _, foodName := range gptRes {
			foodImageDTO := CreateV1RecommendFoodImageDTO(e, foodName)
			foodImage, err := d.Repository.FindOneOrCreateFoodImage(ctx, foodImageDTO)
			if err != nil {
				return response.ResV1RecommendFood{}, err
			}
			foodDTO := CreateV1RecommendFoodDTO(e, foodName, int(foodImage.ID))

			foods, err := d.Repository.SaveRecommendFood(ctx, foodDTO)
			if err != nil {
				return response.ResV1RecommendFood{}, err
			}
			food := response.V1RecommendFood{
				Name: foods.Name,
			}
			imageUrl, err := aws.ImageGetSignedURL(ctx, foodImage.Image, aws.ImgTypeFood)
			if err != nil {
				return response.ResV1RecommendFood{}, err
			}
			food.Image = imageUrl
			aiRes.FoodNames = append(aiRes.FoodNames, food)
			break
		}
		return aiRes, nil
	} else {
		res := response.ResV1RecommendFood{}
		query += " ORDER BY RAND() LIMIT 1"
		food, err := d.Repository.FindOneV1RecommendFood(ctx, query)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		ResFood := response.V1RecommendFood{
			Name: food.Name,
		}
		//food image ID로 이미지 URL을 가져온다.
		image, err := d.Repository.FindOneFoodImage(ctx, food.FoodImageID)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		imageUrl, err := aws.ImageGetSignedURL(ctx, image, aws.ImgTypeFood)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		ResFood.Image = imageUrl
		res.FoodNames = append(res.FoodNames, ResFood)
		return res, nil
	}
}
