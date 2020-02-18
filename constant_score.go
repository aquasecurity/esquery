package esquery

import "encoding/json"

type ConstantScoreQuery struct {
	params constantScoreParams
}

type constantScoreParams struct {
	Filter json.Marshaler `json:"filter"`
	Boost  float32        `json:"boost,omitempty"`
}

func ConstantScore(filter json.Marshaler) *ConstantScoreQuery {
	return &ConstantScoreQuery{
		params: constantScoreParams{Filter: filter},
	}
}

func (q *ConstantScoreQuery) Boost(b float32) *ConstantScoreQuery {
	q.params.Boost = b
	return q
}

func (q ConstantScoreQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]constantScoreParams{
		"constant_score": q.params,
	})
}
