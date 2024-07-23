package response

type ResMetaData struct {
	MetaData MetaData `json:"metaData"`
}

type MetaData struct {
	Types     []string `json:"types"`
	Scenarios []string `json:"scenarios"`
	Times     []string `json:"times"`
}
