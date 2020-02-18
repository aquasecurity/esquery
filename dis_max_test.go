package esquery

import (
	"testing"
)

func TestDisMax(t *testing.T) {
	runTests(t, []queryTest{
		{
			"dis_max",
			DisMax(Term("title", "Quick pets"), Term("body", "Quick pets")).TieBreaker(0.7),
			"{\"dis_max\":{\"queries\":[{\"term\":{\"title\":{\"value\":\"Quick pets\"}}},{\"term\":{\"body\":{\"value\":\"Quick pets\"}}}],\"tie_breaker\":0.7}}\n",
		},
	})
}
