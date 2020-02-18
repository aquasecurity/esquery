package esquery

import "encoding/json"

/*******************************************************************************
 * Boosting Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html
 ******************************************************************************/

type BoostingQuery struct {
	params boostingQueryParams
}

type boostingQueryParams struct {
	Positive      json.Marshaler `json:"positive"`
	Negative      json.Marshaler `json:"negative"`
	NegativeBoost float32        `json:"negative_boost"`
}

func Boosting() *BoostingQuery {
	return &BoostingQuery{}
}

func (q *BoostingQuery) Positive(p json.Marshaler) *BoostingQuery {
	q.params.Positive = p
	return q
}

func (q *BoostingQuery) Negative(p json.Marshaler) *BoostingQuery {
	q.params.Negative = p
	return q
}

func (q *BoostingQuery) NegativeBoost(b float32) *BoostingQuery {
	q.params.NegativeBoost = b
	return q
}

func (q *BoostingQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]boostingQueryParams{
		"boosting": q.params,
	})
}
