package esquery

/*******************************************************************************
 * Boosting Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html
 ******************************************************************************/

type BoostingQuery struct {
	Pos      Mappable
	Neg      Mappable
	NegBoost float32
}

func Boosting() *BoostingQuery {
	return &BoostingQuery{}
}

func (q *BoostingQuery) Positive(p Mappable) *BoostingQuery {
	q.Pos = p
	return q
}

func (q *BoostingQuery) Negative(p Mappable) *BoostingQuery {
	q.Neg = p
	return q
}

func (q *BoostingQuery) NegativeBoost(b float32) *BoostingQuery {
	q.NegBoost = b
	return q
}

func (q *BoostingQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"boosting": map[string]interface{}{
			"positive":       q.Pos.Map(),
			"negative":       q.Neg.Map(),
			"negative_boost": q.NegBoost,
		},
	}
}
