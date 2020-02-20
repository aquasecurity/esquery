package esquery

import (
	"github.com/fatih/structs"
)

/*******************************************************************************
 * Exists Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
 ******************************************************************************/

type ExistsQuery struct {
	Field string `structs:"field"`
}

func Exists(field string) *ExistsQuery {
	return &ExistsQuery{field}
}

func (q *ExistsQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"exists": structs.Map(q),
	}
}

/*******************************************************************************
 * IDs Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-ids-query.html
 ******************************************************************************/

type IDsQuery struct {
	IDs struct {
		Values []string `structs:"values"`
	} `structs:"ids"`
}

func IDs(vals ...string) *IDsQuery {
	q := &IDsQuery{}
	q.IDs.Values = vals
	return q
}

func (q *IDsQuery) Map() map[string]interface{} {
	return structs.Map(q)
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
	Value   string `structs:"value"`
	Rewrite string `structs:"rewrite,omitempty"`
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

func (q *PrefixQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"prefix": map[string]interface{}{
			q.field: structs.Map(q.params),
		},
	}
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
	Gt       interface{}   `structs:"gt,omitempty"`
	Gte      interface{}   `structs:"gte,omitempty"`
	Lt       interface{}   `structs:"lt,omitempty"`
	Lte      interface{}   `structs:"lte,omitempty"`
	Format   string        `structs:"format,omitempty"`
	Relation RangeRelation `structs:"relation,string,omitempty"`
	TimeZone string        `structs:"time_zone,omitempty"`
	Boost    float32       `structs:"boost,omitempty"`
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

func (a *RangeQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"range": map[string]interface{}{
			a.field: structs.Map(a.params),
		},
	}
}

type RangeRelation uint8

const (
	INTERSECTS RangeRelation = iota
	CONTAINS
	WITHIN
)

func (a RangeRelation) String() string {
	switch a {
	case INTERSECTS:
		return "INTERSECTS"
	case CONTAINS:
		return "CONTAINS"
	case WITHIN:
		return "WITHIN"
	default:
		return ""
	}
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
	Value                 string `structs:"value"`
	Flags                 string `structs:"flags,omitempty"`
	MaxDeterminizedStates uint16 `structs:"max_determinized_states,omitempty"`
	Rewrite               string `structs:"rewrite,omitempty"`
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

func (q *RegexpQuery) Map() map[string]interface{} {
	var qType string
	if q.wildcard {
		qType = "wildcard"
	} else {
		qType = "regexp"
	}
	return map[string]interface{}{
		qType: map[string]interface{}{
			q.field: structs.Map(q.params),
		},
	}
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
	Value          string `structs:"value"`
	Fuzziness      string `structs:"fuzziness,omitempty"`
	MaxExpansions  uint16 `structs:"max_expansions,omitempty"`
	PrefixLength   uint16 `structs:"prefix_length,omitempty"`
	Transpositions *bool  `structs:"transpositions,omitempty"`
	Rewrite        string `structs:"rewrite,omitempty"`
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

func (q *FuzzyQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"fuzzy": map[string]interface{}{
			q.field: structs.Map(q.params),
		},
	}
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
	Value interface{} `structs:"value"`
	Boost float32     `structs:"boost,omitempty"`
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

func (q *TermQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"term": map[string]interface{}{
			q.field: structs.Map(q.params),
		},
	}
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

func (q TermsQuery) Map() map[string]interface{} {
	innerMap := map[string]interface{}{q.field: q.values}
	if q.boost > 0 {
		innerMap["boost"] = q.boost
	}

	return map[string]interface{}{"terms": innerMap}
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
	Terms                    []string `structs:"terms"`
	MinimumShouldMatchField  string   `structs:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript string   `structs:"minimum_should_match_script,omitempty"`
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

func (q TermsSetQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"terms_set": map[string]interface{}{
			q.field: structs.Map(q.params),
		},
	}
}
