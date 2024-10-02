package request

type ReqRecommendFood struct {
	Types          string `json:"types" example:"한식"`
	Scenarios      string `json:"scenarios" example:"혼밥"`
	Times          string `json:"times" example:"점심"`
	Themes         string `json:"themes" example:"스트레스 해소"`
	Flavors        string `json:"flavors" example:"매운맛"`
	PreviousAnswer string `json:"previousAnswer" example:"김치찌개 떡볶이 치킨"`
}
