package schema

import "zinccli/meta"

type IndexRequest struct {
	Name        string        `json:"name"`
	StorageType string        `json:"storage_type"`
	Mappings    meta.Mappings `json:"mappings"`
}

type IndexResponse struct {
	Index   string `json:"index"`
	Message string `json:"message"`
	Storage string `json:"storage"`
}
