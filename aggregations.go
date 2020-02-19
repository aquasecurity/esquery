package esquery

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Aggregation interface {
	Mappable
	Name() string
}

type AggregationRequest struct {
	Aggs map[string]Mappable
}

func Aggregate(aggs ...Aggregation) *AggregationRequest {
	req := &AggregationRequest{
		Aggs: make(map[string]Mappable),
	}
	for _, agg := range aggs {
		req.Aggs[agg.Name()] = agg
	}

	return req
}

func (req *AggregationRequest) Map() map[string]interface{} {
	m := make(map[string]interface{})

	for name, agg := range req.Aggs {
		m[name] = agg.Map()
	}

	return map[string]interface{}{
		"aggs": m,
	}
}

func (req *AggregationRequest) Run(
	api *elasticsearch.Client,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*esapi.SearchRequest){api.Search.WithBody(&b)}, o...)

	return api.Search(opts...)
}
