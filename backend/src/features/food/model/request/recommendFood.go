package request

type ReqRecommendFood struct {
	Type     []string `json:"type"`
	Scenario []string `json:"scenario"`
	Time     []string `json:"time"`
}
