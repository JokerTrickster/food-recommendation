package usecase

import (
	"context"
	"fmt"
	"main/features/food/model/entity"
	_errors "main/features/food/model/errors"
	"main/features/food/model/response"
	"main/utils"
	"main/utils/aws"
	"strings"

	"github.com/sashabaranov/go-openai"

	_interface "main/features/food/model/interface"
	"time"
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
		// 음식 추천 로직 구현
		client := openai.NewClient(utils.OpenAIKey)
		if client == nil {
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), "OpenAI Client 초기화 실패", utils.ErrFromChatGPT)
		}

		// 데이터 가공
		question := CreateV1RecommendFoodQuestion(e)
		// ChatGPT 요청 메시지 생성
		prompt := `
너는 한국 요리에 대한 깊은 지식을 가진 전문가이며, 음식 이름과 영양 정보를 제공할 수 있는 AI이다.
한국에 살고 있는 사람들에게 상황에 맞는 음식을 추천해주는 전문가로서, 음식 이름을 추천하기 위해 아래 조건이 있다. 

1. 질문에 적합한 음식 이름을 반드시 1개만 추천한다.
2. 음식 이름과 함께 그 음식의 대략적인 영양 정보(용량(g), 칼로리(kcal), 탄수화물(g), 단백질(g), 지방(g))를 제공해야 한다.
3. 음식 이름을 제외한 다른 정보(요리법, 재료, 가게 이름, 날짜 등)로 대답하면 안된다.
4. 음식 이름에는 공백이 있으면 안된다.
5. 응답에는 이모티콘이나 특수문자(*&^$@~!@...)를 포함하지 않는다.

예시:
- 질문: 오늘 친구와 같이 점심에 매운 음식을 먹고 싶어. 무엇을 추천해줄래?
음식 이름, 용량, 칼로리, 탄수화물, 단백질, 지방 순으로 아래와 같이 응답해줘 단, 음식이름에는 공백이 있으면 안된다.
비빔밥 500g 750 150.01 20.22 10.12

질문을 할 테니, 위와 같은 형식으로 대답해줘:
` + question

		// ChatGPT API 호출
		resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // 사용할 모델 지정
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
			},
		})
		if err != nil {
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), "ChatGPT 호출 실패: "+err.Error(), utils.ErrFromChatGPT)
		}

		// ChatGPT 응답 처리
		gptRes := make([]string, 0)
		if len(resp.Choices) > 0 {
			choice := resp.Choices[0].Message.Content
			cleanedString := strings.Trim(choice, "[] \n")
			gptRes = SplitAndRemoveEmpty(cleanedString)
		} else {
			return response.ResV1RecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrFoodNotFound.Error(), utils.ErrFromChatGPT)
		}
		aiRes := response.ResV1RecommendFood{}

		foodName, nutrition, err := ParseFoodResponse(gptRes)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		fmt.Println(foodName, nutrition)

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
		nutrientDTO, err := d.Repository.FindOneAndSaveNutrient(ctx, nutrition)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		food := response.V1RecommendFood{
			Name:         foods.Name,
			Amount:       nutrientDTO.Amount,
			Kcal:         nutrientDTO.Kcal,
			Carbohydrate: nutrientDTO.Carbohydrate,
			Protein:      nutrientDTO.Protein,
			Fat:          nutrientDTO.Fat,
		}
		imageUrl, err := aws.ImageGetSignedURL(ctx, foodImage.Image, aws.ImgTypeFood)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		food.Image = imageUrl
		aiRes.FoodNames = append(aiRes.FoodNames, food)
		return aiRes, nil
	} else {
		query += " ORDER BY RAND() LIMIT 1"
		food, err := d.Repository.FindOneV1RecommendFood(ctx, query)
		if err != nil {
			return response.ResV1RecommendFood{}, err
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
		nutrientDTO, err := d.Repository.FindOneNutrient(ctx, food.Name)
		if err != nil {
			return response.ResV1RecommendFood{}, err
		}
		res := CreateRes1Recommend(food, imageUrl, nutrientDTO)
		return res, nil
	}
}
