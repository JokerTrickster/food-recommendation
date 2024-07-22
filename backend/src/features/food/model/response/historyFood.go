package response

type ResHistoryFood struct {
	Foods []HistoryFood `json:"foods"`
}

type HistoryFood struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Time     string `json:"time"`
	Scenario string `json:"scenario"`
	Created  string `json:"created"`
}
