package response

type ResMetaData struct {
	MetaKeys []string `json:"metaKeys"`
	MetaData MetaData `json:"metaData"`
}

type MetaData struct {
	Types     []Category `json:"types"`
	Scenarios []Category `json:"scenarios"`
	Times     []Category `json:"times"`
	Themes    []Category `json:"themes"`
	Flavors   []Category `json:"flavors"`
}

type Category struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
