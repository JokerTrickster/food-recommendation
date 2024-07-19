package entity

type RecommendFoodEntity struct {
	Type     string `json:"type"`
	Scenario string `json:"scenario"`
	Time     string `json:"time"`
	UserID   uint   `json:"userID"`
}
