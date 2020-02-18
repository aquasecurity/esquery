package esquery

import (
	"testing"
)

func TestBool(t *testing.T) {
	runTests(t, []queryTest{
		{
			"bool with only a simple must",
			Bool().Must(Term("tag", "tech")),
			"{\"bool\":{\"must\":[{\"term\":{\"tag\":{\"value\":\"tech\"}}}]}}\n",
		},
		{
			"bool which must match_all and filter",
			Bool().Must(MatchAll()).Filter(Term("status", "active")),
			"{\"bool\":{\"must\":[{\"match_all\":{}}],\"filter\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}\n",
		},
		{
			"bool with a lot of stuff",
			Bool().
				Must(Term("user", "kimchy")).
				Filter(Term("tag", "tech")).
				MustNot(Range("age").Gte(10).Lte(20)).
				Should(Term("tag", "wow"), Term("tag", "elasticsearch")).
				MinimumShouldMatch(1).
				Boost(1.1),
			"{\"bool\":{\"must\":[{\"term\":{\"user\":{\"value\":\"kimchy\"}}}],\"filter\":[{\"term\":{\"tag\":{\"value\":\"tech\"}}}],\"must_not\":[{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}],\"should\":[{\"term\":{\"tag\":{\"value\":\"wow\"}}},{\"term\":{\"tag\":{\"value\":\"elasticsearch\"}}}],\"minimum_should_match\":1,\"boost\":1.1}}\n",
		},
	})
}
