package esquery

import "github.com/fatih/structs"

/*******************************************************************************
 * Disjunction Max Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html
 ******************************************************************************/

type DisMaxQuery struct {
	queries    []Mappable
	tieBreaker float32
}

func DisMax(queries ...Mappable) *DisMaxQuery {
	return &DisMaxQuery{
		queries: queries,
	}
}

func (q *DisMaxQuery) TieBreaker(b float32) *DisMaxQuery {
	q.tieBreaker = b
	return q
}

func (q *DisMaxQuery) Map() map[string]interface{} {
	inner := make([]map[string]interface{}, len(q.queries))
	for i, iq := range q.queries {
		inner[i] = iq.Map()
	}
	return map[string]interface{}{
		"dis_max": structs.Map(struct {
			Queries    []map[string]interface{} `structs:"queries"`
			TieBreaker float32                  `structs:"tie_breaker,omitempty"`
		}{inner, q.tieBreaker}),
	}
}
