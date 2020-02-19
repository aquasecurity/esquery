package esquery

import (
	"testing"
)

func TestTermLevel(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"exists",
			Exists("title"),
			map[string]interface{}{
				"exists": map[string]interface{}{
					"field": "title",
				},
			},
		},
		{
			"ids",
			IDs("1", "4", "100"),
			map[string]interface{}{
				"ids": map[string]interface{}{
					"values": []string{"1", "4", "100"},
				},
			},
		},
		{
			"simple prefix",
			Prefix("user", "ki"),
			map[string]interface{}{
				"prefix": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "ki",
					},
				},
			},
		},
		{
			"complex prefix",
			Prefix("user", "ki").Rewrite("ji"),
			map[string]interface{}{
				"prefix": map[string]interface{}{
					"user": map[string]interface{}{
						"value":   "ki",
						"rewrite": "ji",
					},
				},
			},
		},
		{
			"int range",
			Range("age").Gte(10).Lte(20).Boost(2.0),
			map[string]interface{}{
				"range": map[string]interface{}{
					"age": map[string]interface{}{
						"gte":   10,
						"lte":   20,
						"boost": 2.0,
					},
				},
			},
		},
		{
			"string range",
			Range("timestamp").Gte("now-1d/d").Lt("now/d").Relation(CONTAINS),
			map[string]interface{}{
				"range": map[string]interface{}{
					"timestamp": map[string]interface{}{
						"gte":      "now-1d/d",
						"lt":       "now/d",
						"relation": "CONTAINS",
					},
				},
			},
		},
		{
			"regexp",
			Regexp("user", "k.*y").Flags("ALL").MaxDeterminizedStates(10000).Rewrite("constant_score"),
			map[string]interface{}{
				"regexp": map[string]interface{}{
					"user": map[string]interface{}{
						"value":                   "k.*y",
						"flags":                   "ALL",
						"max_determinized_states": 10000,
						"rewrite":                 "constant_score",
					},
				},
			},
		},
		{
			"wildcard",
			Wildcard("user", "ki*y").Rewrite("constant_score"),
			map[string]interface{}{
				"wildcard": map[string]interface{}{
					"user": map[string]interface{}{
						"value":   "ki*y",
						"rewrite": "constant_score",
					},
				},
			},
		},
		{
			"fuzzy",
			Fuzzy("user", "ki").Fuzziness("AUTO").MaxExpansions(50).Transpositions(true),
			map[string]interface{}{
				"fuzzy": map[string]interface{}{
					"user": map[string]interface{}{
						"value":          "ki",
						"fuzziness":      "AUTO",
						"max_expansions": 50,
						"transpositions": true,
					},
				},
			},
		},
		{
			"term",
			Term("user", "Kimchy").Boost(1.3),
			map[string]interface{}{
				"term": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "Kimchy",
						"boost": 1.3,
					},
				},
			},
		},
		{
			"terms",
			Terms("user").Values("bla", "pl").Boost(1.3),
			map[string]interface{}{
				"terms": map[string]interface{}{
					"user":  []string{"bla", "pl"},
					"boost": 1.3,
				},
			},
		},
		{
			"terms_set",
			TermsSet("programming_languages", "go", "rust", "COBOL").MinimumShouldMatchField("required_matches"),
			map[string]interface{}{
				"terms_set": map[string]interface{}{
					"programming_languages": map[string]interface{}{
						"terms":                      []string{"go", "rust", "COBOL"},
						"minimum_should_match_field": "required_matches",
					},
				},
			},
		},
	})
}
