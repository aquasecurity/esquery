package esquery

import (
	"testing"
)

func TestAggregations(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple, single aggregation",
			Aggregate(
				Avg("average_score", "score"),
			),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"average_score": map[string]interface{}{
						"avg": map[string]interface{}{
							"field": "score",
						},
					},
				},
			},
		},
		{
			"a complex, multi-aggregation",
			Aggregate(
				Sum("total_score", "score"),
				WeightedAvg("weighted_score").
					Value("score", 50).
					Weight("weight", 1),
				StringStats("tag_stats", "tags").ShowDistribution(true),
			),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"total_score": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "score",
						},
					},
					"weighted_score": map[string]interface{}{
						"weighted_avg": map[string]interface{}{
							"value": map[string]interface{}{
								"field":   "score",
								"missing": 50,
							},
							"weight": map[string]interface{}{
								"field":   "weight",
								"missing": 1,
							},
						},
					},
					"tag_stats": map[string]interface{}{
						"string_stats": map[string]interface{}{
							"field":             "tags",
							"show_distribution": true,
						},
					},
				},
			},
		},
		{
			"a complex, multi-aggregation, nested",
			Aggregate(
				NestedAgg("categories", "categories").
					Aggs(TermsAgg("type", "outdoors")),
				FilterAgg("filtered",
					Term("type", "t-shirt")),
			),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"categories": map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "categories",
						},
						"aggs": map[string]interface{}{
							"type": map[string]interface{}{
								"terms": map[string]interface{}{
									"field": "outdoors",
								},
							},
						},
					},
					"filtered": map[string]interface{}{
						"filter": map[string]interface{}{
							"term": map[string]interface{}
								"type":   map[string]interface{} {

									"value": "t-shirt",
								},
							},
						},
					},
				},
			},
		},

		{
			"order for termsAggs",
			//eq.Aggregate(eq.TermsAgg("a1", "FIELD1").Size(0).Aggs(eq.Sum("a2", "FIELD2.SUBFIELD")))
			Aggregate(
				TermsAgg("categories", "categories").
					Order(map[string]string{"priceSum": "desc"}).
					Size(5).Aggs(Sum("priceSum", "price"))),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"categories": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "categories",
							"order": map[string]interface{}{
								"priceSum": "desc",
							},
							"size": 5,
						},
						"aggs": map[string]interface{}{
							"priceSum": map[string]interface{}{
								"sum": map[string]interface{}{
									"field": "price",
								},
							},
						},
					},
				},
			},
		},
		{
			"Single include for termsAggs",
			//eq.Aggregate(eq.TermsAgg("a1", "FIELD1").Size(0).Aggs(eq.Sum("a2", "FIELD2.SUBFIELD")))
			Aggregate(
				TermsAgg("categories", "categories").
					Include("red.*|blue.*"),
			),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"categories": map[string]interface{}{
						"terms": map[string]interface{}{
							"field":   "categories",
							"include": "red.*|blue.*",
						},
					},
				},
			},
		},
		{
			"Multi include for termsAggs",
			//eq.Aggregate(eq.TermsAgg("a1", "FIELD1").Size(0).Aggs(eq.Sum("a2", "FIELD2.SUBFIELD")))
			Aggregate(
				TermsAgg("categories", "categories").
					Include("red", "blue"),
			),
			map[string]interface{}{
				"aggs": map[string]interface{}{
					"categories": map[string]interface{}{
						"terms": map[string]interface{}{
							"field":   "categories",
							"include": []string{"red", "blue"},
						},
					},
				},
			},
		},
	})
}
