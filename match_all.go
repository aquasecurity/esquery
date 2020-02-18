package esquery

import (
	"encoding/json"
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
