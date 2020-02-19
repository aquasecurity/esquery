package esquery

import "testing"

func TestMetricAggs(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"avg agg: simple",
			Avg("average_score", "score"),
			map[string]interface{}{
				"avg": map[string]interface{}{
					"field": "score",
				},
			},
		},
		{
			"avg agg: with missing",
			Avg("average_score", "score").Missing(2),
			map[string]interface{}{
				"avg": map[string]interface{}{
					"field":   "score",
					"missing": 2,
				},
			},
		},
		{
			"weighted avg",
			WeightedAvg("weighted_grade").Value("grade", 2).Weight("weight"),
			map[string]interface{}{
				"weighted_avg": map[string]interface{}{
					"value": map[string]interface{}{
						"field":   "grade",
						"missing": 2,
					},
					"weight": map[string]interface{}{
						"field": "weight",
					},
				},
			},
		},
		{
			"cardinality: no precision threshold",
			Cardinality("type_count", "type"),
			map[string]interface{}{
				"cardinality": map[string]interface{}{
					"field": "type",
				},
			},
		},
		{
			"cardinality: with precision threshold",
			Cardinality("type_count", "type").PrecisionThreshold(100),
			map[string]interface{}{
				"cardinality": map[string]interface{}{
					"field":               "type",
					"precision_threshold": 100,
				},
			},
		},
		{
			"value_count agg: simple",
			ValueCount("num_values", "score"),
			map[string]interface{}{
				"value_count": map[string]interface{}{
					"field": "score",
				},
			},
		},
		{
			"sum agg: simple",
			Sum("total_score", "score").Missing(1),
			map[string]interface{}{
				"sum": map[string]interface{}{
					"field":   "score",
					"missing": 1,
				},
			},
		},
		{
			"max agg: simple",
			Max("max_score", "score"),
			map[string]interface{}{
				"max": map[string]interface{}{
					"field": "score",
				},
			},
		},
		{
			"min agg: simple",
			Min("min_score", "score"),
			map[string]interface{}{
				"min": map[string]interface{}{
					"field": "score",
				},
			},
		},
		{
			"percentiles: simple",
			Percentiles("load_time_outlier", "load_time"),
			map[string]interface{}{
				"percentiles": map[string]interface{}{
					"field": "load_time",
				},
			},
		},
		{
			"percentiles: complex",
			Percentiles("load_time_outlier", "load_time").
				Keyed(true).
				Percents(95, 99, 99.9).
				Compression(200).
				NumHistogramDigits(3).
				Missing(20),
			map[string]interface{}{
				"percentiles": map[string]interface{}{
					"field":    "load_time",
					"percents": []float32{95, 99, 99.9},
					"keyed":    true,
					"missing":  20,
					"tdigest": map[string]interface{}{
						"compression": 200,
					},
					"hdr": map[string]interface{}{
						"number_of_significant_value_digits": 3,
					},
				},
			},
		},
		{
			"stats agg",
			Stats("grades_stats", "grade"),
			map[string]interface{}{
				"stats": map[string]interface{}{
					"field": "grade",
				},
			},
		},
		{
			"string_stats agg: no show distribution",
			StringStats("message_stats", "message.keyword"),
			map[string]interface{}{
				"string_stats": map[string]interface{}{
					"field": "message.keyword",
				},
			},
		},
		{
			"string_stats agg: with show distribution",
			StringStats("message_stats", "message.keyword").ShowDistribution(false),
			map[string]interface{}{
				"string_stats": map[string]interface{}{
					"field":             "message.keyword",
					"show_distribution": false,
				},
			},
		},
	})
}
