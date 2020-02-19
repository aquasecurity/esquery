package esquery

import "testing"

func TestCustomAgg(t *testing.T) {
	m := map[string]interface{}{
		"genres": map[string]interface{}{
			"terms": map[string]interface{}{
				"field": "genre",
			},
			"t_shirts": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"type": "t-shirt",
					},
				},
				"aggs": map[string]interface{}{
					"avg_price": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "price",
						},
					},
				},
			},
		},
	}

	runMapTests(t, []mapTest{
		{
			"custom aggregation",
			CustomAgg(m),
			m,
		},
	})
}
