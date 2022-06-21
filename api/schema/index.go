package schema

import "zinccli/meta"

type CreateIndexRequest struct {
	Name        string        `json:"name"`
	StorageType string        `json:"storage_type"`
	Mappings    meta.Mappings `json:"mappings"`
}

type IndexMsgResponse struct {
	Index   string `json:"index"`
	Message string `json:"message"`
	Storage string `json:"storage"`
}

type IndexResponse meta.Index

type MappingRespone meta.Mappings

type MappingMsgRespone struct {
	Message string `json:"message"`
}
