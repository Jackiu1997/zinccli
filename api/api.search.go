package api

import (
	"encoding/json"
	"zinccli/api/schema"
)

func (z *ZincApi) Search(indexName string, query string) (*schema.SearchResponse, error) {
	url := "/" + indexName + "/_search"

	resp, err := z.Request("POST", url, query)
	if err != nil {
		return nil, err
	}

	var searchResp schema.SearchResponse
	err = json.Unmarshal(resp, &searchResp)
	if err != nil {
		return nil, err
	}
	return &searchResp, nil
}
