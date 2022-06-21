package api

import (
	"encoding/json"
	"zinccli/api/schema"
)

func (z *ZincApi) CreateDocument(indexName string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc"

	resp, err := z.Request("POST", url, data)
	if err != nil {
		return nil, err
	}

	var docResp schema.DocumentResponse
	err = json.Unmarshal(resp, &docResp)
	if err != nil {
		return nil, err
	}
	return &docResp, nil
}

func (z *ZincApi) CreateDocumentWithId(indexName string, id string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc/" + id

	resp, err := z.Request("POST", url, data)
	if err != nil {
		return nil, err
	}

	var docResp schema.DocumentResponse
	err = json.Unmarshal(resp, &docResp)
	if err != nil {
		return nil, err
	}
	return &docResp, nil
}

func (z *ZincApi) UpdateDocument(indexName string, id string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_update/" + id

	resp, err := z.Request("POST", url, data)
	if err != nil {
		return nil, err
	}

	var docResp schema.DocumentResponse
	err = json.Unmarshal(resp, &docResp)
	if err != nil {
		return nil, err
	}
	return &docResp, nil
}

func (z *ZincApi) DeleteDocument(indexName string, id string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc/" + id

	resp, err := z.Request("DELETE", url, "")
	if err != nil {
		return nil, err
	}

	var docResp schema.DocumentResponse
	err = json.Unmarshal(resp, &docResp)
	if err != nil {
		return nil, err
	}
	return &docResp, nil
}

func (z *ZincApi) Bulk(data string) (*schema.BulkResponse, error) {
	url := "/_bulk"

	resp, err := z.Request("POST", url, data)
	if err != nil {
		return nil, err
	}

	var bulkResp schema.BulkResponse
	err = json.Unmarshal(resp, &bulkResp)
	if err != nil {
		return nil, err
	}
	return &bulkResp, nil
}
