package esquery

import (
	"bytes"
	"io"

	"github.com/fatih/structs"
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

func (a *MatchQuery) Map() map[string]interface{} {
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

	return map[string]interface{}{
		mType: map[string]interface{}{
			a.field: structs.Map(a.params),
		},
	}
}

type matchParams struct {
	Qry          interface{}   `structs:"query"`
	Anl          string        `structs:"analyzer,omitempty"`
	AutoGenerate *bool         `structs:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzz         string        `structs:"fuzziness,omitempty"`
	MaxExp       uint16        `structs:"max_expansions,omitempty"`
	PrefLen      uint16        `structs:"prefix_length,omitempty"`
	Trans        *bool         `structs:"transpositions,omitempty"`
	FuzzyRw      string        `structs:"fuzzy_rewrite,omitempty"`
	Lent         bool          `structs:"lenient,omitempty"`
	Op           MatchOperator `structs:"operator,string,omitempty"`
	MinMatch     string        `structs:"minimum_should_match,omitempty"`
	ZeroTerms    ZeroTerms     `structs:"zero_terms_query,string,omitempty"`
	Slp          uint16        `structs:"slop,omitempty"` // only relevant for match_phrase query
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
	if len(simpleQuery) > 0 {
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

func (q *MatchQuery) ZeroTermsQuery(s ZeroTerms) *MatchQuery {
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

func (a MatchOperator) String() string {
	switch a {
	case OR:
		return "or"
	case AND:
		return "and"
	default:
		return ""
	}
}

type ZeroTerms uint8

const (
	None ZeroTerms = iota
	All
)

func (a ZeroTerms) String() string {
	switch a {
	case None:
		return "none"
	case All:
		return "all"
	default:
		return ""
	}
}
