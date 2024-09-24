package entity

type SelectFoodEntity struct {
	Type     string `json:"type"`
	Scenario string `json:"scenario"`
	Time     string `json:"time"`
	Theme    string `json:"theme"`
	Flavor   string `json:"flavor"`
	Name     string `json:"name"`
	UserID   uint   `json:"userID"`
}
