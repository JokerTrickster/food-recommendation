package request

type ReqSelectFood struct {
	Type     string `json:"type" example:"한식"`     // 전체, 양식, 한식, 중식 등
	Scenario string `json:"scenario" example:"혼밥"` // 전체 , 혼밥, 가족, 친구들
	Time     string `json:"time" example:"점심"`
	Theme    string `json:"theme" example:"스트레스 해소"`
	Flavor   string `json:"flavor" example:"매운맛"`
	Name     string `json:"name" example:"된장찌개"`
}
