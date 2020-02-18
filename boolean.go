package esquery

import "encoding/json"

/*******************************************************************************
 * Boolean Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
 ******************************************************************************/

type BoolQuery struct {
	params boolQueryParams
}

type boolQueryParams struct {
	Must               []json.Marshaler `json:"must,omitempty"`
	Filter             []json.Marshaler `json:"filter,omitempty"`
	MustNot            []json.Marshaler `json:"must_not,omitempty"`
	Should             []json.Marshaler `json:"should,omitempty"`
	MinimumShouldMatch int16            `json:"minimum_should_match,omitempty"`
	Boost              float32          `json:"boost,omitempty"`
}

func Bool() *BoolQuery {
	return &BoolQuery{}
}

func (q *BoolQuery) Must(must ...json.Marshaler) *BoolQuery {
	q.params.Must = append(q.params.Must, must...)
	return q
}

func (q *BoolQuery) Filter(filter ...json.Marshaler) *BoolQuery {
	q.params.Filter = append(q.params.Filter, filter...)
	return q
}

func (q *BoolQuery) MustNot(mustnot ...json.Marshaler) *BoolQuery {
	q.params.MustNot = append(q.params.MustNot, mustnot...)
	return q
}

func (q *BoolQuery) Should(should ...json.Marshaler) *BoolQuery {
	q.params.Should = append(q.params.Should, should...)
	return q
}

func (q *BoolQuery) MinimumShouldMatch(val int16) *BoolQuery {
	q.params.MinimumShouldMatch = val
	return q
}

func (q *BoolQuery) Boost(val float32) *BoolQuery {
	q.params.Boost = val
	return q
}

func (q BoolQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]boolQueryParams{
		"bool": q.params,
	})
}
