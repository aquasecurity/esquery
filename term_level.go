package esquery

import (
	"encoding/json"
)

/*******************************************************************************
 * Exists Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
 ******************************************************************************/

type ExistsQuery string

func Exists(field string) *ExistsQuery {
	q := ExistsQuery(field)
	return &q
}

func (q ExistsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"exists": map[string]string{
			"field": string(q),
		},
	})
}

/*******************************************************************************
 * IDs Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html
 ******************************************************************************/

type IDsQuery []string

func IDs(vals ...string) *IDsQuery {
	q := IDsQuery(vals)
	return &q
}

func (q IDsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"ids": map[string][]string{
			"values": []string(q),
		},
	})
}

/*******************************************************************************
 * Prefix Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
 ******************************************************************************/

type PrefixQuery struct {
	field  string
	params prefixQueryParams
}

type prefixQueryParams struct {
	Value   string `json:"value"`
	Rewrite string `json:"rewrite,omitempty"`
}

func Prefix(field, value string) *PrefixQuery {
	return &PrefixQuery{
		field:  field,
		params: prefixQueryParams{Value: value},
	}
}

func (q *PrefixQuery) Rewrite(s string) *PrefixQuery {
	q.params.Rewrite = s
	return q
}

func (q PrefixQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"prefix": map[string]prefixQueryParams{
			q.field: q.params,
		},
	})
}

/*******************************************************************************
 * Range Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html
 ******************************************************************************/

type RangeQuery struct {
	field  string
	params rangeQueryParams
}

type rangeQueryParams struct {
	Gt       interface{}   `json:"gt,omitempty"`
	Gte      interface{}   `json:"gte,omitempty"`
	Lt       interface{}   `json:"lt,omitempty"`
	Lte      interface{}   `json:"lte,omitempty"`
	Format   string        `json:"format,omitempty"`
	Relation RangeRelation `json:"relation,omitempty"`
	TimeZone string        `json:"time_zone,omitempty"`
	Boost    float32       `json:"boost,omitempty"`
}

func Range(field string) *RangeQuery {
	return &RangeQuery{field: field}
}

func (a *RangeQuery) Gt(val interface{}) *RangeQuery {
	a.params.Gt = val
	return a
}

func (a *RangeQuery) Gte(val interface{}) *RangeQuery {
	a.params.Gte = val
	return a
}

func (a *RangeQuery) Lt(val interface{}) *RangeQuery {
	a.params.Lt = val
	return a
}

func (a *RangeQuery) Lte(val interface{}) *RangeQuery {
	a.params.Lte = val
	return a
}

func (a *RangeQuery) Format(f string) *RangeQuery {
	a.params.Format = f
	return a
}

func (a *RangeQuery) Relation(r RangeRelation) *RangeQuery {
	a.params.Relation = r
	return a
}

func (a *RangeQuery) TimeZone(zone string) *RangeQuery {
	a.params.TimeZone = zone
	return a
}

func (a *RangeQuery) Boost(b float32) *RangeQuery {
	a.params.Boost = b
	return a
}

func (a RangeQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"range": map[string]rangeQueryParams{
			a.field: a.params,
		},
	})
}

type RangeRelation uint8

const (
	INTERSECTS RangeRelation = iota
	CONTAINS
	WITHIN
)

func (a RangeRelation) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case INTERSECTS:
		s = "INTERSECTS"
	case CONTAINS:
		s = "CONTAINS"
	case WITHIN:
		s = "WITHIN"
	default:
		return nil, ErrInvalidValue
	}

	return json.Marshal(s)
}

/*******************************************************************************
 * Regexp Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html
 ******************************************************************************/

type RegexpQuery struct {
	field    string
	wildcard bool
	params   regexpQueryParams
}

type regexpQueryParams struct {
	Value                 string `json:"value"`
	Flags                 string `json:"flags,omitempty"`
	MaxDeterminizedStates uint16 `json:"max_determinized_states,omitempty"`
	Rewrite               string `json:"rewrite,omitempty"`
}

func Regexp(field, value string) *RegexpQuery {
	return &RegexpQuery{
		field: field,
		params: regexpQueryParams{
			Value: value,
		},
	}
}

func (q *RegexpQuery) Value(v string) *RegexpQuery {
	q.params.Value = v
	return q
}

func (q *RegexpQuery) Flags(f string) *RegexpQuery {
	if !q.wildcard {
		q.params.Flags = f
	}
	return q
}

func (q *RegexpQuery) MaxDeterminizedStates(m uint16) *RegexpQuery {
	if !q.wildcard {
		q.params.MaxDeterminizedStates = m
	}
	return q
}

func (q *RegexpQuery) Rewrite(r string) *RegexpQuery {
	q.params.Rewrite = r
	return q
}

func (q RegexpQuery) MarshalJSON() ([]byte, error) {
	var qType string
	if q.wildcard {
		qType = "wildcard"
	} else {
		qType = "regexp"
	}
	return json.Marshal(map[string]interface{}{
		qType: map[string]regexpQueryParams{
			q.field: q.params,
		},
	})
}

/*******************************************************************************
 * Wildcard Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
 ******************************************************************************/

func Wildcard(field, value string) *RegexpQuery {
	return &RegexpQuery{
		field:    field,
		wildcard: true,
		params: regexpQueryParams{
			Value: value,
		},
	}
}

/*******************************************************************************
 * Fuzzy Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
 ******************************************************************************/

type FuzzyQuery struct {
	field  string
	params fuzzyQueryParams
}

type fuzzyQueryParams struct {
	Value          string `json:"value"`
	Fuzziness      string `json:"fuzziness,omitempty"`
	MaxExpansions  uint16 `json:"max_expansions,omitempty"`
	PrefixLength   uint16 `json:"prefix_length,omitempty"`
	Transpositions *bool  `json:"transpositions,omitempty"`
	Rewrite        string `json:"rewrite,omitempty"`
}

func Fuzzy(field, value string) *FuzzyQuery {
	return &FuzzyQuery{
		field: field,
		params: fuzzyQueryParams{
			Value: value,
		},
	}
}

func (q *FuzzyQuery) Value(val string) *FuzzyQuery {
	q.params.Value = val
	return q
}

func (q *FuzzyQuery) Fuzziness(fuzz string) *FuzzyQuery {
	q.params.Fuzziness = fuzz
	return q
}

func (q *FuzzyQuery) MaxExpansions(m uint16) *FuzzyQuery {
	q.params.MaxExpansions = m
	return q
}

func (q *FuzzyQuery) PrefixLength(l uint16) *FuzzyQuery {
	q.params.PrefixLength = l
	return q
}

func (q *FuzzyQuery) Transpositions(b bool) *FuzzyQuery {
	q.params.Transpositions = &b
	return q
}

func (q *FuzzyQuery) Rewrite(s string) *FuzzyQuery {
	q.params.Rewrite = s
	return q
}

func (q FuzzyQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"fuzzy": map[string]fuzzyQueryParams{
			q.field: q.params,
		},
	})
}

/*******************************************************************************
 * Term Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
 ******************************************************************************/

type TermQuery struct {
	field  string
	params termQueryParams
}

type termQueryParams struct {
	Value interface{} `json:"value"`
	Boost float32     `json:"boost,omitempty"`
}

func Term(field string, value interface{}) *TermQuery {
	return &TermQuery{
		field: field,
		params: termQueryParams{
			Value: value,
		},
	}
}

func (q *TermQuery) Value(val interface{}) *TermQuery {
	q.params.Value = val
	return q
}

func (q *TermQuery) Boost(b float32) *TermQuery {
	q.params.Boost = b
	return q
}

func (q TermQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"term": map[string]termQueryParams{
			q.field: q.params,
		},
	})
}

/*******************************************************************************
 * Terms Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-query.html
 ******************************************************************************/

type TermsQuery struct {
	field  string
	values []interface{}
	boost  float32
}

func Terms(field string, values ...interface{}) *TermsQuery {
	return &TermsQuery{
		field:  field,
		values: values,
	}
}

func (q *TermsQuery) Values(values ...interface{}) *TermsQuery {
	q.values = values
	return q
}

func (q *TermsQuery) Boost(b float32) *TermsQuery {
	q.boost = b
	return q
}

func (q TermsQuery) MarshalJSON() ([]byte, error) {
	innerMap := map[string]interface{}{q.field: q.values}
	if q.boost > 0 {
		innerMap["boost"] = q.boost
	}
	return json.Marshal(map[string]interface{}{"terms": innerMap})
}

/*******************************************************************************
 * Term Set Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-terms-set-query.html
 ******************************************************************************/

type TermsSetQuery struct {
	field  string
	params termsSetQueryParams
}

type termsSetQueryParams struct {
	Terms                    []string `json:"terms"`
	MinimumShouldMatchField  string   `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript string   `json:"minimum_should_match_script,omitempty"`
}

func TermsSet(field string, terms ...string) *TermsSetQuery {
	return &TermsSetQuery{
		field: field,
		params: termsSetQueryParams{
			Terms: terms,
		},
	}
}

func (q *TermsSetQuery) Terms(terms ...string) *TermsSetQuery {
	q.params.Terms = terms
	return q
}

func (q *TermsSetQuery) MinimumShouldMatchField(field string) *TermsSetQuery {
	q.params.MinimumShouldMatchField = field
	return q
}

func (q *TermsSetQuery) MinimumShouldMatchScript(script string) *TermsSetQuery {
	q.params.MinimumShouldMatchScript = script
	return q
}

func (q TermsSetQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"terms_set": map[string]termsSetQueryParams{
			q.field: q.params,
		},
	})
}
