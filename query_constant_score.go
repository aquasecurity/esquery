package esquery

import "github.com/fatih/structs"

/*******************************************************************************
 * Constant Score Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html
 ******************************************************************************/

type ConstantScoreQuery struct {
	filter Mappable
	boost  float32
}

func ConstantScore(filter Mappable) *ConstantScoreQuery {
	return &ConstantScoreQuery{
		filter: filter,
	}
}

func (q *ConstantScoreQuery) Boost(b float32) *ConstantScoreQuery {
	q.boost = b
	return q
}

func (q *ConstantScoreQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"constant_score": structs.Map(struct {
			Filter map[string]interface{} `structs:"filter"`
			Boost  float32                `structs:"boost,omitempty"`
		}{q.filter.Map(), q.boost}),
	}
}
