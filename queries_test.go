package esquery

import (
	"testing"
)

func TestQueries(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple match_all query",
			Query(MatchAll()),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
			},
		},
		{
			"a complex query",
			Query(
				Bool().
					Must(
						Range("date").
							Gt("some time in the past").
							Lte("now").
							Relation(CONTAINS).
							TimeZone("Asia/Jerusalem").
							Boost(2.3),

						Match("author").
							Query("some guy").
							Analyzer("analyzer?").
							Fuzziness("fuzz"),
					).
					Boost(3.1),
			),
			map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []map[string]interface{}{
							{
								"range": map[string]interface{}{
									"date": map[string]interface{}{
										"gt":        "some time in the past",
										"lte":       "now",
										"relation":  "CONTAINS",
										"time_zone": "Asia/Jerusalem",
										"boost":     2.3,
									},
								},
							},
							{
								"match": map[string]interface{}{
									"author": map[string]interface{}{
										"query":     "some guy",
										"analyzer":  "analyzer?",
										"fuzziness": "fuzz",
									},
								},
							},
						},
						"boost": 3.1,
					},
				},
			},
		},
	})
}
