package esquery

import (
	"bytes"
	"encoding/json"
	"testing"
)

type queryTest struct {
	name    string
	q       json.Marshaler
	expJSON string
}

func runTests(t *testing.T, tests []queryTest) {
	for _, test := range tests {
		var b bytes.Buffer
		t.Run(test.name, func(t *testing.T) {
			err := encode(test.q, &b)
			if err != nil {
				t.Errorf("unexpectedly failed: %s", err)
			} else if b.String() != test.expJSON {
				t.Errorf("expected %q, got %q", test.expJSON, b.String())
			}
		})
	}
}
