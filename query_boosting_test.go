package esquery

import (
	"testing"
)

func TestBoosting(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"boosting query",
			Boosting().
				Positive(Term("text", "apple")).
				Negative(Term("text", "pie tart")).
				NegativeBoost(0.5),
			map[string]interface{}{
				"boosting": map[string]interface{}{
					"positive": map[string]interface{}{
						"term": map[string]interface{}{
							"text": map[string]interface{}{
								"value": "apple",
							},
						},
					},
					"negative": map[string]interface{}{
						"term": map[string]interface{}{
							"text": map[string]interface{}{
								"value": "pie tart",
							},
						},
					},
					"negative_boost": 0.5,
				},
			},
		},
	})
}
