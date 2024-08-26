package entity

type SelectFoodEntity struct {
	Type     string `json:"type"`
	Scenario string `json:"scenario"`
	Time     string `json:"time"`
	Name     string `json:"name"`
	UserID   uint   `json:"userID"`
}
