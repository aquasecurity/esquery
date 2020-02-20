package esquery

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type QueryRequest struct {
	Query Mappable
}

func Query(q Mappable) *QueryRequest {
	return &QueryRequest{q}
}

func (req *QueryRequest) Map() map[string]interface{} {
	return map[string]interface{}{
		"query": req.Query.Map(),
	}
}

func (req *QueryRequest) Run(
	api *elasticsearch.Client,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Query.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*esapi.SearchRequest){api.Search.WithBody(&b)}, o...)

	return api.Search(opts...)
}
