package request

type ReqRecommendFood struct {
	Type           string `json:"type" example:"한식"`
	Scenario       string `json:"scenario" example:"혼밥"`
	Time           string `json:"time" example:"점심"`
	Theme          string `json:"theme" example:"스트레스 해소"`
	Flavor         string `json:"flavor" example:"매운맛"`
	PreviousAnswer string `json:"previousAnswer" example:"김치찌개 떡볶이 치킨"`
}
