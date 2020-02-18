package esquery

import (
	"testing"
)

func TestTermLevel(t *testing.T) {
	runTests(t, []queryTest{
		{"exists", Exists("title"), "{\"exists\":{\"field\":\"title\"}}\n"},

		{"ids", IDs("1", "4", "100"), "{\"ids\":{\"values\":[\"1\",\"4\",\"100\"]}}\n"},

		{"simple prefix", Prefix("user", "ki"), "{\"prefix\":{\"user\":{\"value\":\"ki\"}}}\n"},

		{"complex prefix", Prefix("user", "ki").Rewrite("ji"), "{\"prefix\":{\"user\":{\"value\":\"ki\",\"rewrite\":\"ji\"}}}\n"},

		{"int range", Range("age").Gte(10).Lte(20).Boost(2.0), "{\"range\":{\"age\":{\"gte\":10,\"lte\":20,\"boost\":2}}}\n"},

		{"string range", Range("timestamp").Gte("now-1d/d").Lt("now/d").Relation(CONTAINS), "{\"range\":{\"timestamp\":{\"gte\":\"now-1d/d\",\"lt\":\"now/d\",\"relation\":\"CONTAINS\"}}}\n"},

		{"regexp", Regexp("user", "k.*y").Flags("ALL").MaxDeterminizedStates(10000).Rewrite("constant_score"), "{\"regexp\":{\"user\":{\"value\":\"k.*y\",\"flags\":\"ALL\",\"max_determinized_states\":10000,\"rewrite\":\"constant_score\"}}}\n"},

		{"wildcard", Wildcard("user", "ki*y").Rewrite("constant_score"), "{\"wildcard\":{\"user\":{\"value\":\"ki*y\",\"rewrite\":\"constant_score\"}}}\n"},

		{"fuzzy", Fuzzy("user", "ki").Fuzziness("AUTO").MaxExpansions(50).Transpositions(true), "{\"fuzzy\":{\"user\":{\"value\":\"ki\",\"fuzziness\":\"AUTO\",\"max_expansions\":50,\"transpositions\":true}}}\n"},

		{"term", Term("user", "Kimchy").Boost(1.3), "{\"term\":{\"user\":{\"value\":\"Kimchy\",\"boost\":1.3}}}\n"},

		{"terms", Terms("user").Values("bla", "pl").Boost(1.3), "{\"terms\":{\"boost\":1.3,\"user\":[\"bla\",\"pl\"]}}\n"},

		{"terms_set", TermsSet("programming_languages", "go", "rust", "COBOL").MinimumShouldMatchField("required_matches"), "{\"terms_set\":{\"programming_languages\":{\"terms\":[\"go\",\"rust\",\"COBOL\"],\"minimum_should_match_field\":\"required_matches\"}}}\n"},
	})
}
