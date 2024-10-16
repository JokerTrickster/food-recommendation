package request

type ReqSaveFood struct {
	Foods []SaveFood `json:"foods"`
}

type SaveFood struct {
	Name      string `json:"name" validate:"required"`
	Times     string `json:"times"`
	Scenarios string `json:"scenarios"`
	Flavors   string `json:"flavors"`
	Themes    string `json:"themes"`
	Types     string `json:"types"`
}
