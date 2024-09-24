package entity

type RecommendFoodEntity struct {
	Type           string `json:"type"`
	Scenario       string `json:"scenario"`
	Time           string `json:"time"`
	Theme          string `json:"theme"`
	Flavor         string `json:"flavor"`
	UserID         uint   `json:"userID"`
	PreviousAnswer string `json:"previousAnswer"`
}

type Content struct {
	Parts []string `json:"Parts"`
	Role  string   `json:"Role"`
}
type Candidates struct {
	Content *Content `json:"Content"`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:"Candidates"`
}
