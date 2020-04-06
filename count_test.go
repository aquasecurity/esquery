package esquery

import "testing"

func TestCount(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple count request",
			Count(MatchAll()),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
	})
}
