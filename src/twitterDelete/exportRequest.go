package twitterDelete

// ExportRequest -- request representing export to either CSV or JSON file
type ExportRequest struct {
	Username    string      `json:"userName"`
	Extension   string      `json:"extension"`
	Credentials Credentials `jsong:"credentials"`
}
