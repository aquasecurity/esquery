package esquery

import "testing"

func TestCustomQuery(t *testing.T) {
	m := map[string]interface{}{
		"geo_distance": map[string]interface{}{
			"distance": "200km",
			"pin.location": map[string]interface{}{
				"lat": 40,
				"lon": -70,
			},
		},
	}

	runMapTests(t, []mapTest{
		{
			"custom query",
			CustomQuery(m),
			m,
		},
	})
}
