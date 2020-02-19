package esquery

import (
	"testing"
)

func TestBool(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"bool with only a simple must",
			Bool().Must(Term("tag", "tech")),
			map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"tag": map[string]interface{}{
									"value": "tech",
								},
							},
						},
					},
				},
			},
		},
		{
			"bool which must match_all and filter",
			Bool().Must(MatchAll()).Filter(Term("status", "active")),
			map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{"match_all": map[string]interface{}{}},
					},
					"filter": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"status": map[string]interface{}{
									"value": "active",
								},
							},
						},
					},
				},
			},
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
			map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"user": map[string]interface{}{
									"value": "kimchy",
								},
							},
						},
					},
					"filter": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"tag": map[string]interface{}{
									"value": "tech",
								},
							},
						},
					},
					"must_not": []map[string]interface{}{
						{
							"range": map[string]interface{}{
								"age": map[string]interface{}{
									"gte": 10,
									"lte": 20,
								},
							},
						},
					},
					"should": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"tag": map[string]interface{}{
									"value": "wow",
								},
							},
						},
						{
							"term": map[string]interface{}{
								"tag": map[string]interface{}{
									"value": "elasticsearch",
								},
							},
						},
					},
					"minimum_should_match": 1,
					"boost":                1.1,
				},
			},
		},
	})
}
