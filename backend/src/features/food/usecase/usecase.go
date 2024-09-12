package usecase

import (
	"context"
	"fmt"
	"main/features/food/model/entity"
	"main/features/food/model/response"
	"main/utils/aws"
	"main/utils/db/mysql"
	"strings"
	"time"
)

func CreateDailyRecommendFoodQuestion() string {
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("오늘 날짜 %s와 궁합이 좋은 음식 3개 추천해줘", today)
	return question
}
func CreateResEmptyImageFood(foods []mysql.Foods) response.ResEmptyImageFood {
	var res response.ResEmptyImageFood
	for _, f := range foods {
		emptyFood := response.EmptyFood{
			ID:   f.ID,
			Name: f.Name,
		}
		res.Foods = append(res.Foods, emptyFood)
	}
	return res
}
func CreateSelectFoodQuestion(e entity.SelectFoodEntity) string {
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("오늘 날짜 %s 와 %s 음식 궁합을 알려줘", today, e.Name)
	return question
}
func CreateResRankingFood(foodList []string) response.ResRankingFood {
	res := response.ResRankingFood{}
	res.Foods = foodList
	return res
}

func CreateResMetaData(typeDTO []mysql.Types, timeDTO []mysql.Times, scenarioDTO []mysql.Scenarios, themesDTO []mysql.Themes, flavorDTO []mysql.Flavors) response.ResMetaData {
	var res response.ResMetaData
	var metaData response.MetaData
	for _, t := range typeDTO {
		category := response.Category{
			Name: t.Name,
		}
		imageUrl, err := aws.ImageGetSignedURL(context.TODO(), t.Image, aws.ImgTypeCategory)
		if err != nil {
			return response.ResMetaData{}
		}
		category.Image = imageUrl
		metaData.Types = append(metaData.Types, category)
	}
	for _, t := range timeDTO {
		category := response.Category{
			Name:  t.Name,
			Image: t.Image,
		}
		imageUrl, err := aws.ImageGetSignedURL(context.TODO(), t.Image, aws.ImgTypeCategory)
		if err != nil {
			return response.ResMetaData{}
		}
		category.Image = imageUrl
		metaData.Times = append(metaData.Times, category)
	}
	for _, t := range scenarioDTO {
		category := response.Category{
			Name:  t.Name,
			Image: t.Image,
		}
		imageUrl, err := aws.ImageGetSignedURL(context.TODO(), t.Image, aws.ImgTypeCategory)
		if err != nil {
			return response.ResMetaData{}
		}
		category.Image = imageUrl
		metaData.Scenarios = append(metaData.Scenarios, category)
	}
	for _, t := range themesDTO {
		category := response.Category{
			Name:  t.Name,
			Image: t.Image,
		}
		imageUrl, err := aws.ImageGetSignedURL(context.TODO(), t.Image, aws.ImgTypeCategory)
		if err != nil {
			return response.ResMetaData{}
		}
		category.Image = imageUrl
		metaData.Themes = append(metaData.Themes, category)
	}
	for _, t := range flavorDTO {
		category := response.Category{
			Name:  t.Name,
			Image: t.Image,
		}
		imageUrl, err := aws.ImageGetSignedURL(context.TODO(), t.Image, aws.ImgTypeCategory)
		if err != nil {
			return response.ResMetaData{}
		}
		category.Image = imageUrl
		metaData.Flavors = append(metaData.Flavors, category)
	}
	res.MetaData = metaData
	return res
}

func CreateSelectFoodDTO(e entity.SelectFoodEntity) *mysql.Foods {
	typeID, err := mysql.GetTypeID(e.Type)
	if err != nil {
		fmt.Println(err)
	}
	timeID, err := mysql.GetTimeID(e.Time)
	if err != nil {
		fmt.Println(err)
	}
	secnarioID, err := mysql.GetScenarioID(e.Scenario)
	if err != nil {
		fmt.Println(err)
	}
	return &mysql.Foods{
		TypeID:     typeID,
		TimeID:     timeID,
		ScenarioID: secnarioID,
		Name:       e.Name,
	}
}
func CreateFoodHistoryDTO(foodID, userID uint) *mysql.FoodHistory {
	return &mysql.FoodHistory{
		FoodID: foodID,
		UserID: userID,
	}
}

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
	themeID, err := mysql.GetThemeID(entity.Theme)
	if err != nil {
		fmt.Println(err)
	}
	flavorID, err := mysql.GetFlavorID(entity.Flavor)
	if err != nil {
		fmt.Println(err)
	}

	return &mysql.Foods{
		TypeID:     typeID,
		TimeID:     timeID,
		ScenarioID: secnarioID,
		ThemeID:    themeID,
		FlavorID:   flavorID,
		Name:       foodName,
		Image:      "food_default.png",
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
	if entity.Type == "" {
		reqType = "전체 음식"
	} else {
		reqType = entity.Type
	}
	var reqScenario string
	if entity.Scenario == "" {
		reqScenario = "누구든지"
	} else {
		reqScenario = entity.Scenario
	}
	var reqTime string
	if entity.Time == "" {
		reqTime = "아무때나"
	} else {
		reqTime = entity.Time
	}
	var reqTheme string
	if entity.Theme == "" {
		reqTheme = "아무 테마"
	} else {
		reqTheme = entity.Theme
	}
	var reqFlavor string
	if entity.Flavor == "" {
		reqFlavor = "아무맛"
	} else {
		reqFlavor = entity.Flavor
	}

	questionType := fmt.Sprintf("어떤 종류의 음식 :  %s \n", reqType)
	questionScenario := fmt.Sprintf("누구와 함께 : %s \n", reqScenario)
	questionTime := fmt.Sprintf("언제 : %s \n", reqTime)
	questionTheme := fmt.Sprintf("어떤 테마 : %s \n", reqTheme)
	questionFlavor := fmt.Sprintf("어떤 맛 : %s \n", reqFlavor)
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("%s와 어울리는 %s, %s, %s, %s, %s, 음식 이름 1개만 추천해줘 설명 필요없고 이름만 추천해줘", today, questionType, questionScenario, questionTime, questionTheme, questionFlavor)
	if entity.PreviousAnswer != "" {
		question += fmt.Sprintf("이전에 추천받은 음식은 제외하고 알려줘 이전 추천 음식 이름 : %s", entity.PreviousAnswer)
	}

	return question
}
