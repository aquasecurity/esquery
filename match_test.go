package esquery

import (
	"testing"
)

func TestMatch(t *testing.T) {
	runTests(t, []queryTest{
		{"simple match", Match("title", "sample text"), "{\"match\":{\"title\":{\"query\":\"sample text\"}}}\n"},
		{"match with more params", Match("issue_number").Query(16).Transpositions(false).MaxExpansions(32).Operator(AND), "{\"match\":{\"issue_number\":{\"query\":16,\"max_expansions\":32,\"transpositions\":false,\"operator\":\"and\"}}}\n"},
		{"match_bool_prefix", MatchBoolPrefix("title", "sample text"), "{\"match_bool_prefix\":{\"title\":{\"query\":\"sample text\"}}}\n"},
		{"match_phrase", MatchPhrase("title", "sample text"), "{\"match_phrase\":{\"title\":{\"query\":\"sample text\"}}}\n"},
		{"match_phrase_prefix", MatchPhrasePrefix("title", "sample text"), "{\"match_phrase_prefix\":{\"title\":{\"query\":\"sample text\"}}}\n"},
	})
}
