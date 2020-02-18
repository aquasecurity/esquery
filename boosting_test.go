package esquery

import (
	"testing"
)

func TestBoost(t *testing.T) {
	runTests(t, []queryTest{
		{
			"boosting query",
			Boosting().
				Positive(Term("text", "apple")).
				Negative(Term("text", "pie tart")).
				NegativeBoost(0.5),
			"{\"boosting\":{\"positive\":{\"term\":{\"text\":{\"value\":\"apple\"}}},\"negative\":{\"term\":{\"text\":{\"value\":\"pie tart\"}}},\"negative_boost\":0.5}}\n",
		},
	})
}
