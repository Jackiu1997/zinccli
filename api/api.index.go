package api

import (
	"encoding/json"
	"zinccli/api/schema"
)

func (z *ZincApi) CreateIndex(indexName string, data *schema.CreateIndexRequest) (*schema.IndexResponse, error) {
	url := "/index/" + indexName

	query, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := z.Request("POST", url, string(query))
	if err != nil {
		return nil, err
	}

	var indexResponse schema.IndexResponse
	err = json.Unmarshal(resp, &indexResponse)
	if err != nil {
		return nil, err
	}
	return &indexResponse, nil
}

func (z *ZincApi) DeleteIndex(indexName string) (*schema.IndexResponse, error) {
	url := "/index/" + indexName

	resp, err := z.Request("DELETE", url, "")
	if err != nil {
		return nil, err
	}

	var indexResponse schema.IndexResponse
	err = json.Unmarshal(resp, &indexResponse)
	if err != nil {
		return nil, err
	}
	return &indexResponse, nil
}

func (z *ZincApi) GetMapping(indexName string) (*schema.MappingRespone, error) {
	url := "/" + indexName + "_mapping"

	resp, err := z.Request("GET", url, "")
	if err != nil {
		return nil, err
	}

	// convert response to map
	var m = make(map[string]interface{})
	err = json.Unmarshal(resp, &m)
	if err != nil {
		return nil, err
	}

	// extarct mappings body
	var m1 = make(map[string]interface{})
	err = json.Unmarshal(m[indexName].([]byte), &m1)
	if err != nil {
		return nil, err
	}

	var mappingRespone schema.MappingRespone
	err = json.Unmarshal(m1["mappings"].([]byte), &mappingRespone)
	if err != nil {
		return nil, err
	}
	return &mappingRespone, nil
}

func (z *ZincApi) UpdateMapping(indexName string, data string) (*schema.MappingMsgRespone, error) {
	url := "/" + indexName + "_mapping"

	resp, err := z.Request("PUT", url, data)
	if err != nil {
		return nil, err
	}

	var updateMappingRespone schema.MappingMsgRespone
	err = json.Unmarshal(resp, &updateMappingRespone)
	if err != nil {
		return nil, err
	}
	return &updateMappingRespone, nil
}

func (z *ZincApi) ListIndex() ([]*schema.IndexResponse, error) {
	url := "/index"

	resp, err := z.Request("GET", url, "")
	if err != nil {
		return nil, err
	}

	var indexes []schema.IndexResponse
	err = json.Unmarshal(resp, &indexes)
	if err != nil {
		return nil, err
	}

	var res []*schema.IndexResponse
	for _, index := range indexes {
		res = append(res, &index)
	}
	return res, nil
}
