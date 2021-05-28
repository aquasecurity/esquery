package esquery

import (
	"testing"
)

func TestNested(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"Nested Query",
			Nested("dns_values", Term("dns_values.type", "A")).ScoreMode("max").IgnoreUnmapped(true),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path":            "dns_values",
					"query":           Term("dns_values.type", "A").Map(),
					"score_mode":      "max",
					"ignore_unmapped": true,
				},
			},
		},
	})
}
