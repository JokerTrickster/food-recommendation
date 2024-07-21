package usecase

import (
	"fmt"
	"main/features/food/model/entity"
	"main/utils/db/mysql"
	"strings"
)

func CreateRecommendFoodDTO(entity entity.RecommendFoodEntity, foodName string) *mysql.Foods {
	typeID, err := mysql.GetTypeID(entity.Type)
	if err != nil {
		fmt.Println(err)
	}
	timeID, err := mysql.GetTimeID(entity.Time)
	if err != nil {
		fmt.Println(err)
	}
	secnarioID, err := mysql.GetScenarioID(entity.Scenario)
	if err != nil {
		fmt.Println(err)
	}

	return &mysql.Foods{
		TypeID:     typeID,
		TimeID:     timeID,
		ScenarioID: secnarioID,
		Name:       foodName,
	}

}

func SplitAndRemoveEmpty(s string) []string {
	// 문자열의 연속된 공백을 단일 공백으로 치환하고 앞뒤 공백 제거
	trimmedString := strings.TrimSpace(s)
	// 공백을 기준으로 문자열 분할
	words := strings.Fields(trimmedString)
	return words
}

func CreateRecommendFoodQuestion(entity entity.RecommendFoodEntity) string {
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
	if entity.PreviousAnswer != "" {
		question += fmt.Sprintf("이전에 추천받은 음식은 제외하고 알려줘 이전 추천 음식 이름 : %s", entity.PreviousAnswer)
	}

	return question
}
