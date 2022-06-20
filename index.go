package zinccli

import (
	"encoding/json"
	"zinccli/meta"
	"zinccli/schema"
)

func (z *ZincCli) CreateIndex(indexName string, data *schema.IndexRequest) (*schema.IndexResponse, error) {
	url := "/index/" + indexName

	query, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := z.ZincRequest("POST", url, string(query))
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

func (z *ZincCli) DeleteIndex(indexName string) (*schema.IndexResponse, error) {
	url := "/index/" + indexName

	resp, err := z.ZincRequest("DELETE", url, "")
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

func (z *ZincCli) GetMapping(indexName string) (*schema.IndexResponse, error) {
	url := "/" + indexName + "_mapping"

	resp, err := z.ZincRequest("GET", url, "")
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

	var indexResponse schema.IndexResponse
	err = json.Unmarshal(m1["mappings"].([]byte), &indexResponse)
	if err != nil {
		return nil, err
	}
	return &indexResponse, nil
}

func (z *ZincCli) UpdateMapping(indexName string, data string) (*schema.IndexResponse, error) {
	url := "/" + indexName + "_mapping"

	resp, err := z.ZincRequest("PUT", url, data)
	if err != nil {
		return nil, err
	}

	var m = make(map[string]interface{})
	err = json.Unmarshal(resp, &m)
	if err != nil {
		return nil, err
	}
	return &schema.IndexResponse{
		Index:   indexName,
		Message: m["message"].(string),
		Storage: "",
	}, nil
}

func (z *ZincCli) ListIndex() ([]*meta.Index, error) {
	url := "/index"

	resp, err := z.ZincRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	var indexes []meta.Index
	err = json.Unmarshal(resp, &indexes)
	if err != nil {
		return nil, err
	}

	var res []*meta.Index
	for _, index := range indexes {
		res = append(res, &index)
	}
	return res, nil
}
