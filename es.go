package esquery

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
)

type ESQuery struct {
	Query json.Marshaler `json:"query"`
}

func encode(q json.Marshaler, b *bytes.Buffer) (err error) {
	b.Reset()
	err = json.NewEncoder(b).Encode(q)
	if err != nil {
		return fmt.Errorf("failed encoding query to JSON: %w", err)
	}

	return nil
}

func search(q json.Marshaler, api *esapi.API, o ...func(*esapi.SearchRequest)) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = encode(ESQuery{q}, &b)
	if err != nil {
		return res, err
	}

	opts := append([]func(*esapi.SearchRequest){api.Search.WithBody(&b)}, o...)

	return api.Search(opts...)
}

func (q ESQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]json.Marshaler{
		"query": q.Query,
	})
}
