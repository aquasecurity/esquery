package esquery

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// QueryRequest represents a complete request of type "query" to ElasticSearch's
// search API. It simply wraps a value of a type that implements the Mappable
// interface.
type QueryRequest struct {
	Query Mappable
}

// Query generates a search request of type "query", represented by a
// *QueryRequest object. It receives any query type that implements the
// Mappable interface, whether provided internally by the library or custom
// types provided by consuming code.
func Query(q Mappable) *QueryRequest {
	return &QueryRequest{q}
}

// Map implements the Mappable interface. It converts the "query" request into a
// (potentially nested) map[string]interface{}.
func (req *QueryRequest) Map() map[string]interface{} {
	return map[string]interface{}{
		"query": req.Query.Map(),
	}
}

// MarshalJSON implements the json.Marshaler interface, it simply encodes the
// map representation of the query (provided by the Map method) as JSON.
func (req *QueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(req.Map())
}

// Run executes the request using the provided ElasticSearch client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
func (req *QueryRequest) Run(
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
func (req *QueryRequest) RunSearch(
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
