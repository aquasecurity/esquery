package esquery

import (
	"testing"
)

func TestMultiMatch(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"simple multi_match",
			MultiMatch("value1", "value2").Fields("title"),
			map[string]interface{}{
				"multi_match": map[string]interface{}{
					"fields": []string{"title"},
					"query":  "value2",
				},
			},
		},
		{
			"multi_match all params",
			MultiMatch("original").
				Query("test").
				Analyzer("stop").
				Fields("title", "body").
				AutoGenerateSynonymsPhraseQuery(true).
				Fuzziness("AUTO").
				MaxExpansions(16).
				PrefixLength(12).
				TieBreaker(0.3).
				Boost(6.4).
				Transpositions(true).
				FuzzyRewrite("scoring_boolean").
				Lenient(true).
				Operator(OperatorAnd).
				Type(MatchTypePhrase).
				MinimumShouldMatch("3<90%").
				Slop(2).
				ZeroTermsQuery(ZeroTermsAll),
			map[string]interface{}{
				"multi_match": map[string]interface{}{
					"analyzer":                            "stop",
					"auto_generate_synonyms_phrase_query": true,
					"boost":                               6.4,
					"fuzziness":                           "AUTO",
					"fuzzy_rewrite":                       "scoring_boolean",
					"lenient":                             true,
					"max_expansions":                      16,
					"minimum_should_match":                "3<90%",
					"prefix_length":                       12,
					"transpositions":                      true,
					"type":                                "phrase",
					"tie_breaker":                         0.3,
					"operator":                            "AND",
					"zero_terms_query":                    "all",
					"slop":                                2,
					"query":                               "test",
					"fields":                              []string{"title", "body"},
				},
			},
		},
	})
}
