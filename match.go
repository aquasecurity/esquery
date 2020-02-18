package esquery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/esapi"
)

/*******************************************************************************
 * Match Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-bool-prefix-query.html
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase.html
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html
 ******************************************************************************/
type matchType uint8

const (
	TypeMatch matchType = iota
	TypeMatchBoolPrefix
	TypeMatchPhrase
	TypeMatchPhrasePrefix
)

type MatchQuery struct {
	field  string
	mType  matchType
	params matchParams
}

func (a MatchQuery) MarshalJSON() ([]byte, error) {
	var mType string
	switch a.mType {
	case TypeMatch:
		mType = "match"
	case TypeMatchBoolPrefix:
		mType = "match_bool_prefix"
	case TypeMatchPhrase:
		mType = "match_phrase"
	case TypeMatchPhrasePrefix:
		mType = "match_phrase_prefix"
	}

	return json.Marshal(map[string]interface{}{
		mType: map[string]interface{}{
			a.field: a.params,
		},
	})
}

func (a *MatchQuery) Run(api *esapi.API, o ...func(*esapi.SearchRequest)) (res *esapi.Response, err error) {
	return search(*a, api, o...)
}

type matchParams struct {
	Qry          interface{}   `json:"query"`
	Anl          string        `json:"analyzer,omitempty"`
	AutoGenerate *bool         `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzz         string        `json:"fuzziness,omitempty"`
	MaxExp       uint16        `json:"max_expansions,omitempty"`
	PrefLen      uint16        `json:"prefix_length,omitempty"`
	Trans        *bool         `json:"transpositions,omitempty"`
	FuzzyRw      string        `json:"fuzzy_rewrite,omitempty"`
	Lent         bool          `json:"lenient,omitempty"`
	Op           MatchOperator `json:"operator,omitempty"`
	MinMatch     string        `json:"minimum_should_match,omitempty"`
	ZeroTerms    string        `json:"zero_terms_query,omitempty"`
	Slp          uint16        `json:"slop,omitempty"` // only relevant for match_phrase query
}

func Match(fieldName string, simpleQuery ...interface{}) *MatchQuery {
	return newMatch(TypeMatch, fieldName, simpleQuery...)
}

func MatchBoolPrefix(fieldName string, simpleQuery ...interface{}) *MatchQuery {
	return newMatch(TypeMatchBoolPrefix, fieldName, simpleQuery...)
}

func MatchPhrase(fieldName string, simpleQuery ...interface{}) *MatchQuery {
	return newMatch(TypeMatchPhrase, fieldName, simpleQuery...)
}

func MatchPhrasePrefix(fieldName string, simpleQuery ...interface{}) *MatchQuery {
	return newMatch(TypeMatchPhrasePrefix, fieldName, simpleQuery...)
}

func newMatch(mType matchType, fieldName string, simpleQuery ...interface{}) *MatchQuery {
	var qry interface{}
	if simpleQuery != nil && len(simpleQuery) > 0 {
		qry = simpleQuery[len(simpleQuery)-1]
	}

	return &MatchQuery{
		field: fieldName,
		mType: mType,
		params: matchParams{
			Qry: qry,
		},
	}
}

func (q *MatchQuery) Query(data interface{}) *MatchQuery {
	q.params.Qry = data
	return q
}

func (q *MatchQuery) Analyzer(a string) *MatchQuery {
	q.params.Anl = a
	return q
}

func (q *MatchQuery) AutoGenerateSynonymsPhraseQuery(b bool) *MatchQuery {
	q.params.AutoGenerate = &b
	return q
}

func (q *MatchQuery) Fuzziness(f string) *MatchQuery {
	q.params.Fuzz = f
	return q
}

func (q *MatchQuery) MaxExpansions(e uint16) *MatchQuery {
	q.params.MaxExp = e
	return q
}

func (q *MatchQuery) PrefixLength(l uint16) *MatchQuery {
	q.params.PrefLen = l
	return q
}

func (q *MatchQuery) Transpositions(b bool) *MatchQuery {
	q.params.Trans = &b
	return q
}

func (q *MatchQuery) FuzzyRewrite(s string) *MatchQuery {
	q.params.FuzzyRw = s
	return q
}

func (q *MatchQuery) Lenient(b bool) *MatchQuery {
	q.params.Lent = b
	return q
}

func (q *MatchQuery) Operator(op MatchOperator) *MatchQuery {
	q.params.Op = op
	return q
}

func (q *MatchQuery) MinimumShouldMatch(s string) *MatchQuery {
	q.params.MinMatch = s
	return q
}

func (q *MatchQuery) Slop(n uint16) *MatchQuery {
	q.params.Slp = n
	return q
}

func (q *MatchQuery) ZeroTermsQuery(s string) *MatchQuery {
	q.params.ZeroTerms = s
	return q
}

func (q *MatchQuery) Reader() io.Reader {
	var b bytes.Buffer
	return &b
}

type MatchOperator uint8

const (
	OR MatchOperator = iota
	AND
)

var ErrInvalidValue = errors.New("invalid constant value")

func (a MatchOperator) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case OR:
		s = "or"
	case AND:
		s = "and"
	default:
		return nil, ErrInvalidValue
	}

	return json.Marshal(s)
}

type ZeroTerms uint8

const (
	None ZeroTerms = iota
	All
)

func (a ZeroTerms) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case None:
		s = "none"
	case All:
		s = "all"
	default:
		return nil, ErrInvalidValue
	}

	return json.Marshal(s)
}

/*******************************************************************************
 * Multi-Match Queries
 * https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html
 * NOTE: uncommented for now, article is too long
 ******************************************************************************/

//type MultiMatchQuery struct {
//fields []string
//mType  multiMatchType
//params multiMatchQueryParams
//}

//type multiMatchType uint8

//const (
//BestFields multiMatchType = iota
//MostFields
//CrossFields
//Phrase
//PhrasePrefix
//BoolPrefix
//)

//func (a multiMatchType) MarshalJSON() ([]byte, error) {
//var s string
//switch a {
//case BestFields:
//s = "best_fields"
//case MostFields:
//s = "most_fields"
//case CrossFields:
//s = "cross_fields"
//case Phrase:
//s = "phrase"
//case PhrasePrefix:
//s = "phrase_prefix"
//case BoolPrefix:
//s = "bool_prefix"
//default:
//return nil, ErrInvalidValue
//}
//return json.Marshal(s)
//}
