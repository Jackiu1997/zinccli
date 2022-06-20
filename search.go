package zinccli

import (
	"encoding/json"
	"zinccli/schema"
)

func (z *ZincCli) Search(indexName string, query string) (*schema.SearchResponse, error) {
	url := "/" + indexName + "/_search"

	resp, err := z.ZincRequest("POST", url, query)
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
