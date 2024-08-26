package request

type ReqRecommendFood struct {
	Type           string `json:"type" example:"한식"`     // 전체, 양식, 한식, 중식 등
	Scenario       string `json:"scenario" example:"혼밥"` // 전체 , 혼밥, 가족, 친구들
	Time           string `json:"time" example:"중식"`
	PreviousAnswer string `json:"previousAnswer" example:"김치찌개 떡볶이 치킨"`
}
