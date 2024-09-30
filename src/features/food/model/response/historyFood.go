package response

type ResHistoryFood struct {
	Foods []HistoryFood `json:"foods"`
}

type HistoryFood struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Time     string `json:"time"`
	Scenario string `json:"scenario"`
	Theme    string `json:"theme"`
	Flavor   string `json:"flavor"`
	Created  string `json:"created"`
}
