package esquery

import (
	"testing"
)

func TestDisMax(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"dis_max",
			DisMax(Term("title", "Quick pets"), Term("body", "Quick pets")).TieBreaker(0.7),
			map[string]interface{}{
				"dis_max": map[string]interface{}{
					"queries": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"title": map[string]interface{}{
									"value": "Quick pets",
								},
							},
						},
						{
							"term": map[string]interface{}{
								"body": map[string]interface{}{
									"value": "Quick pets",
								},
							},
						},
					},
					"tie_breaker": 0.7,
				},
			},
		},
	})
}
