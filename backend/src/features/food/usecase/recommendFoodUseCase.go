package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/features/food/model/entity"
	_interface "main/features/food/model/interface"
	"main/utils"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Content struct {
	Parts []string `json:Parts`
	Role  string   `json:Role`
}
type Candidates struct {
	Content *Content `json:Content`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:Candidates`
}
type RecommendFoodUseCase struct {
	Repository     _interface.IRecommendFoodRepository
	ContextTimeout time.Duration
}

func NewRecommendFoodUseCase(repo _interface.IRecommendFoodRepository, timeout time.Duration) _interface.IRecommendFoodUseCase {
	return &RecommendFoodUseCase{Repository: repo, ContextTimeout: timeout}
}

func (d *RecommendFoodUseCase) Recommend(c context.Context, entity entity.RecommendFoodEntity) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	// 유저 정보를 가져온다.

	//음식 추천 로직 구현

	client, err := genai.NewClient(ctx, option.WithAPIKey(utils.GeminiID))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	model := client.GenerativeModel("gemini-1.5-flash")

	//데이터 가공
	var reqType string
	if entity.Type == "전체" {
		reqType = "전체 음식"
	} else {
		reqType = entity.Type
	}
	var reqScenario string
	if entity.Scenario == "전체" {
		reqScenario = "누구든지"
	} else {
		reqScenario = entity.Scenario
	}
	var reqTime string
	if entity.Time == "전체" {
		reqTime = "아무때나"
	} else {
		reqTime = entity.Time
	}
	questionType := fmt.Sprintf("어떤 종류의 음식 :  %s \n", reqType)
	questionScenario := fmt.Sprintf("누구와 함께 : %s \n", reqScenario)
	questionTime := fmt.Sprintf("언제 : %s \n", reqTime)
	question := fmt.Sprintf("%s %s %s 음식 이름 4개만 추천해줘 설명 필요없고 이름만 추천해줘", questionType, questionScenario, questionTime)

	resp, err := model.GenerateContent(
		ctx,
		genai.Text("너는 맛있는 요리 음식을 알려주는 전문가이다."),
		genai.Text("내가 질문을 하면 단어로만 대답을 해줘야 된다."),
		genai.Text("예를 들어서 '매운 음식 추천해줘' 라고 물으면 '김치찌개' 라고 대답을 해줘야 된다."),
		genai.Text("여러개라면 '김치찌개 떡볶이 치킨' 이렇게 응답해주면 된다. "),
		genai.Text("지금부터 질문할게 대답해줘"),
		genai.Text(question),
	)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := make([]string, 0)
	// 출력 부분 수정
	if len(resp.Candidates) > 0 {
		marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
		var generateResponse ContentResponse
		if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
			log.Fatal(err)
		}
		for _, cad := range *generateResponse.Candidates {
			if cad.Content != nil {
				cleanedString := strings.Trim(cad.Content.Parts[0], "[] \n")
				res = strings.Split(cleanedString, " ")
			}
		}

	} else {
		fmt.Println("No candidates found in the response")
	}

	//db에 저장

	return res, nil
}
