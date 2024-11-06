package response

type ResV1RecommendFood struct {
	FoodNames []V1RecommendFood `json:"foodNames"`
}

type V1RecommendFood struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
