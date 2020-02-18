package esquery

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/esapi"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html
type MatchAllQuery struct {
	all    bool
	params matchAllParams
}

type matchAllParams struct {
	Boost float32 `json:"boost,omitempty"`
}

func (a MatchAllQuery) MarshalJSON() ([]byte, error) {
	var mType string
	switch a.all {
	case true:
		mType = "match_all"
	default:
		mType = "match_none"
	}

	return json.Marshal(map[string]matchAllParams{mType: a.params})
}

func (a *MatchAllQuery) Run(api *esapi.API, o ...func(*esapi.SearchRequest)) (res *esapi.Response, err error) {
	return search(*a, api, o...)
}

func MatchAll() *MatchAllQuery {
	return &MatchAllQuery{all: true}
}

func (q *MatchAllQuery) Boost(b float32) *MatchAllQuery {
	if q.all {
		q.params.Boost = b
	}
	return q
}

func MatchNone() *MatchAllQuery {
	return &MatchAllQuery{all: false}
}
