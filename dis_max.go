package esquery

import "encoding/json"

type DisMaxQuery struct {
	params disMaxParams
}

type disMaxParams struct {
	Queries    []json.Marshaler `json:"queries"`
	TieBreaker float32          `json:"tie_breaker,omitempty"`
}

func DisMax(queries ...json.Marshaler) *DisMaxQuery {
	return &DisMaxQuery{
		params: disMaxParams{
			Queries: queries,
		},
	}
}

func (q *DisMaxQuery) TieBreaker(b float32) *DisMaxQuery {
	q.params.TieBreaker = b
	return q
}

func (q DisMaxQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]disMaxParams{
		"dis_max": q.params,
	})
}
