package usecase

import (
	"context"
	"main/features/food/model/response"
	"main/utils/aws"

	_interface "main/features/food/model/interface"
	"main/utils"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
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

	//

	// 음식 추천 로직 구현
	client := openai.NewClient(utils.OpenAIKey)
	if client == nil {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), "OpenAI Client 초기화 실패", utils.ErrFromChatGPT)
	}

	// 데이터 가공
	question := CreateDailyRecommendFoodQuestion()

	// 프롬프트 생성
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "너는 한국에 살고 있고 사람들이 많이 먹는 음식 이름을 알고 있는 전문가이다.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "오늘 날짜와 궁합이 좋을 것 같은 음식 이름을 3개 추천해줘.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "예를 들어서 '피자 치킨 탕수육' 이런 식으로 음식 이름 사이에 공백을 추가해서 3개만 대답해주면 된다. 음식 이름 사이에 공백을 추가하지 않으면 1개로 인식한다.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "한식, 중식, 일식, 양식, 패스트 푸드 중 추천해주면 된다.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "지금부터 질문할게 대답해줘:",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: question,
		},
	}

	// ChatCompletion 호출
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	})
	if err != nil {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), err.Error(), utils.ErrFromChatGPT)
	}

	gptRes := make([]string, 0)

	// 응답 처리
	if len(resp.Choices) > 0 {
		content := resp.Choices[0].Message.Content
		cleanedString := strings.Trim(content, "[] \n")
		gptRes = SplitAndRemoveEmpty(cleanedString)
	} else {
		return response.ResDailyRecommendFood{}, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), "응답이 존재하지 않습니다.", utils.ErrFromChatGPT)
	}

	res := response.ResDailyRecommendFood{}
	// DB에서 가져오기
	for i, foodName := range gptRes {
		if i == 3 {
			break
		}
		food := response.DailyRecommendFood{
			Name:  foodName,
			Image: "food_default.png",
		}
		foods, err := d.Repository.FindOneFood(ctx, foodName)
		if err != nil {
			return response.ResDailyRecommendFood{}, err
		}
		if foods != nil {
			foodImage, err := d.Repository.FindOneFoodImage(ctx, foods.FoodImageID)
			if err != nil {
				return response.ResDailyRecommendFood{}, err
			}
			food.Image = foodImage
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
