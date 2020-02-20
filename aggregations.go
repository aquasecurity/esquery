package esquery

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// AggregationRequest represents a complete request of type "aggregations"
// (a.k.a "aggs") to ElasticSearch's search API. It simply wraps a map of named
// aggregations, which are values of a type that implements the Mappable
// interface.
type AggregationRequest struct {
	Aggs map[string]Mappable
}

// Aggregation is an interface that each aggregation type must implement. It
// is simply an extension of the Mappable interface to include a Named function,
// which returns the name of the aggregation.
type Aggregation interface {
	Mappable
	Name() string
}

// Aggregate generates a search request of type "aggs", represented by a
// *AggregationRequest object. It receives a variadic amount of values that
// implement the Aggregation interface, whether provided internally by the
// library or custom aggregations provided by consuming code.
func Aggregate(aggs ...Aggregation) *AggregationRequest {
	req := &AggregationRequest{
		Aggs: make(map[string]Mappable),
	}
	for _, agg := range aggs {
		req.Aggs[agg.Name()] = agg
	}

	return req
}

// Map implements the Mappable interface. It converts the "aggs" request into a
// (potentially nested) map[string]interface{}.
func (req *AggregationRequest) Map() map[string]interface{} {
	m := make(map[string]interface{})

	for name, agg := range req.Aggs {
		m[name] = agg.Map()
	}

	return map[string]interface{}{
		"aggs": m,
	}
}

// MarshalJSON implements the json.Marshaler interface, it simply encodes the
// map representation of the request (provided by the Map method) as JSON.
func (req *AggregationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.Map())
}

// Run executes the request using the provided ElasticSearch client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
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

// RunSearch is the same as the Run method, except that it accepts a value of
// type esapi.Search (usually this is the Search field of an elasticsearch.Client
// object). Since the ElasticSearch client does not provide an interface type
// for its API (which would allow implementation of mock clients), this provides
// a workaround. The Search function in the ES client is actually a field of a
// function type.
func (req *AggregationRequest) RunSearch(
	search esapi.Search,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	opts := append([]func(*esapi.SearchRequest){search.WithBody(&b)}, o...)

	return search(opts...)
}
