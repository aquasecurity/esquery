package esquery

import (
	"testing"
)

func TestMatchAll(t *testing.T) {
	runTests(t, []queryTest{
		{"match_all without a boost", MatchAll(), "{\"match_all\":{}}\n"},
		{"match_all with a boost", MatchAll().Boost(2.3), "{\"match_all\":{\"boost\":2.3}}\n"},
		{"match_none", MatchNone(), "{\"match_none\":{}}\n"},
	})
}
