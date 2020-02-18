package esquery

import (
	"testing"
)

func TestConstantScore(t *testing.T) {
	runTests(t, []queryTest{
		{
			"constant_score query without boost",
			ConstantScore(Term("user", "kimchy")),
			"{\"constant_score\":{\"filter\":{\"term\":{\"user\":{\"value\":\"kimchy\"}}}}}\n",
		},
		{
			"constant_score query with boost",
			ConstantScore(Term("user", "kimchy")).Boost(2.2),
			"{\"constant_score\":{\"filter\":{\"term\":{\"user\":{\"value\":\"kimchy\"}}},\"boost\":2.2}}\n",
		},
	})
}
