package schema

type DocumentResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

type BulkResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}
