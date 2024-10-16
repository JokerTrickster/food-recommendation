package usecase

import (
	"context"
	"fmt"
	"main/features/food/model/entity"
	"main/features/food/model/request"
	"main/features/food/model/response"
	"main/utils/aws"
	"main/utils/db/mysql"
	"strings"
	"time"
)

func CreateRecommendFoodImageDTO(entity entity.RecommendFoodEntity, foodName string) *mysql.FoodImages {

	return &mysql.FoodImages{
		Name:  foodName,
		Image: "food_default.png",
	}
}

func CreateDailyRecommendFoodQuestion() string {
	today := time.Now().Format("2006-01-02")
	question := fmt.Sprintf("오늘 날짜 %s와 궁합이 좋은 음식 3개 추천해줘 음식 이름만 추천해줘", today)
	return question
}
func CreateResEmptyImageFood(foodImages []mysql.FoodImages) response.ResEmptyImageFood {
	var res response.ResEmptyImageFood
	for _, f := range foodImages {
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

func CreateResMetaData(typeDTO []mysql.Types, timeDTO []mysql.Times, scenarioDTO []mysql.Scenarios, themesDTO []mysql.Themes, flavorDTO []mysql.Flavors) response.ResMetaData {
	var res response.ResMetaData
	var metaData response.MetaData
	//상황 -> 시간 -> 종륲별 -> 맛 -> 기분/테마별
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

	res.MetaData = metaData
	res.MetaKeys = []string{"types", "times", "scenarios", "themes", "flavors"}
	return res
}

func CreateSelectFoodDTO(entity entity.SelectFoodEntity) *mysql.Foods {
	var err error
	typeID := 0
	timeID := 0
	secnarioID := 0
	themeID := 0
	flavorID := 0

	if entity.Types != "" {
		typeID, err = mysql.GetTypeID(entity.Types)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Times != "" {
		timeID, err = mysql.GetTimeID(entity.Times)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Scenarios != "" {
		secnarioID, err = mysql.GetScenarioID(entity.Scenarios)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Themes != "" {
		themeID, err = mysql.GetThemeID(entity.Themes)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Flavors != "" {

		flavorID, err = mysql.GetFlavorID(entity.Flavors)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &mysql.Foods{
		TypeID:     typeID,
		TimeID:     timeID,
		ScenarioID: secnarioID,
		ThemeID:    themeID,
		FlavorID:   flavorID,
		Name:       entity.Name,
	}
}
func CreateFoodHistoryDTO(foodID, userID uint, name string) *mysql.FoodHistory {
	return &mysql.FoodHistory{
		FoodID: foodID,
		UserID: userID,
		Name:   name,
	}
}

func CreateRecommendFoodDTO(entity entity.RecommendFoodEntity, foodName string, foodImageID int) *mysql.Foods {
	var err error
	typeID := 0
	timeID := 0
	secnarioID := 0
	themeID := 0
	flavorID := 0

	if entity.Types != "" {
		typeID, err = mysql.GetTypeID(entity.Types)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Times != "" {
		timeID, err = mysql.GetTimeID(entity.Times)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Scenarios != "" {
		secnarioID, err = mysql.GetScenarioID(entity.Scenarios)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Themes != "" {
		themeID, err = mysql.GetThemeID(entity.Themes)
		if err != nil {
			fmt.Println(err)
		}
	}
	if entity.Flavors != "" {

		flavorID, err = mysql.GetFlavorID(entity.Flavors)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &mysql.Foods{
		TypeID:      typeID,
		TimeID:      timeID,
		ScenarioID:  secnarioID,
		ThemeID:     themeID,
		FlavorID:    flavorID,
		Name:        foodName,
		FoodImageID: foodImageID,
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
	if entity.Types == "" || entity.Types == "전체" {
		reqType = "전체 음식"
	} else {
		reqType = entity.Types
	}
	var reqScenario string
	if entity.Scenarios == "" || entity.Scenarios == "전체" {
		reqScenario = "어떤 상황이든"
	} else {
		reqScenario = entity.Scenarios
	}
	var reqTime string
	if entity.Times == "" || entity.Times == "전체" {
		reqTime = "아무때나"
	} else {
		reqTime = entity.Times
	}
	var reqTheme string
	if entity.Themes == "" || entity.Themes == "전체" {
		reqTheme = "아무 테마"
	} else {
		reqTheme = entity.Themes
	}
	var reqFlavor string
	if entity.Flavors == "" || entity.Flavors == "전체" {
		reqFlavor = "모든 맛"
	} else {
		reqFlavor = entity.Flavors
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

func CreateFoodDTOList(req *request.ReqSaveFood) []*mysql.Foods {
	var foods []*mysql.Foods
	for _, f := range req.Foods {
		food := mysql.Foods{
			Name: f.Name,
		}
		foods = append(foods, &food)
	}
	return foods
}

func CreateSaveFoodImageDTO(food request.SaveFood) *mysql.FoodImages {
	return &mysql.FoodImages{
		Name:  food.Name,
		Image: "food_default.png",
	}
}

func CreateSaveFoodDTO(food request.SaveFood, foodImageID int) *mysql.Foods {
	var err error
	typeID := 0
	timeID := 0
	secnarioID := 0
	themeID := 0
	flavorID := 0

	if food.Types != "" {
		typeID, err = mysql.GetTypeID(food.Types)
		if err != nil {
			fmt.Println(err)
		}
	}
	if food.Times != "" {
		timeID, err = mysql.GetTimeID(food.Times)
		if err != nil {
			fmt.Println(err)
		}
	}
	if food.Scenarios != "" {
		secnarioID, err = mysql.GetScenarioID(food.Scenarios)
		if err != nil {
			fmt.Println(err)
		}
	}
	if food.Themes != "" {
		themeID, err = mysql.GetThemeID(food.Themes)
		if err != nil {
			fmt.Println(err)
		}
	}
	if food.Flavors != "" {

		flavorID, err = mysql.GetFlavorID(food.Flavors)
		if err != nil {
			fmt.Println(err)
		}
	}

	return &mysql.Foods{
		TypeID:      typeID,
		TimeID:      timeID,
		ScenarioID:  secnarioID,
		ThemeID:     themeID,
		FlavorID:    flavorID,
		Name:        food.Name,
		FoodImageID: foodImageID,
	}
}
