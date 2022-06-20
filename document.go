package zinccli

import (
	"encoding/json"
	"zinccli/schema"
)

func (z *ZincCli) CreateDocument(indexName string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc"

	resp, err := z.ZincRequest("POST", url, data)
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

func (z *ZincCli) CreateDocumentWithId(indexName string, id string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc/" + id

	resp, err := z.ZincRequest("POST", url, data)
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

func (z *ZincCli) UpdateDocument(indexName string, id string, data string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_update/" + id

	resp, err := z.ZincRequest("POST", url, data)
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

func (z *ZincCli) DeleteDocument(indexName string, id string) (*schema.DocumentResponse, error) {
	url := "/" + indexName + "/_doc/" + id

	resp, err := z.ZincRequest("DELETE", url, "")
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

func (z *ZincCli) Bulk(data string) (*schema.BulkResponse, error) {
	url := "/_bulk"

	resp, err := z.ZincRequest("POST", url, data)
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
