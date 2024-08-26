package usecase

import (
	"context"
	"encoding/json"
	"main/features/food/model/entity"

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

func (d *RecommendFoodUseCase) Recommend(c context.Context, e entity.RecommendFoodEntity) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	//음식 추천 로직 구현
	client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GeminiID))
	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	model := client.GenerativeModel("gemini-1.5-flash")
	//데이터 가공
	question := CreateRecommendFoodQuestion(e)
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("너는 맛있는 요리 음식을 알려주는 전문가이다."),
		genai.Text("내가 질문을 하면 단어로만 대답을 해줘야 된다."),
		genai.Text("예를 들어서 '매운 음식 추천해줘' 라고 물으면 '김치찌개' 라고 대답을 해줘야 된다."),
		genai.Text("반드시 음식 이름만 추천해줘"),
		genai.Text("여러개라면 '김치찌개 떡볶이 치킨' 이렇게 응답해주면 된다. "),
		genai.Text("지금부터 질문할게 대답해줘"),
		genai.Text(question),
	)

	if err != nil {
		return nil, utils.ErrorMsg(ctx, utils.ErrPartner, utils.Trace(), _errors.ErrGeminiError.Error()+err.Error(), utils.ErrFromGemini)
	}
	res := make([]string, 0)
	// 출력 부분 수정

	if len(resp.Candidates) > 0 {
		marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
		var generateResponse entity.ContentResponse
		if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
			return nil, utils.ErrorMsg(ctx, utils.ErrInternalServer, utils.Trace(), _errors.ErrServerError.Error()+err.Error(), utils.ErrFromInternal)
		}
		for _, cad := range *generateResponse.Candidates {
			if cad.Content != nil {
				cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
				res = SplitAndRemoveEmpty(cleanedString)
			}
		}

	} else {
		return nil, utils.ErrorMsg(ctx, utils.ErrNotFound, utils.Trace(), _errors.ErrFoodNotFound.Error(), utils.ErrFromGemini)
	}

	//db에 저장
	for _, foodName := range res {
		foodDTO := CreateRecommendFoodDTO(e, foodName)
		if err := d.Repository.SaveRecommendFood(ctx, foodDTO); err != nil {
			return nil, err
		}
	}

	return res, nil
}
