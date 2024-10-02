package entity

type SelectFoodEntity struct {
	Types     string `json:"types"`
	Scenarios string `json:"scenarios"`
	Times     string `json:"times"`
	Themes    string `json:"themes"`
	Flavors   string `json:"flavors"`
	Name      string `json:"name"`
	UserID    uint   `json:"userID"`
}
