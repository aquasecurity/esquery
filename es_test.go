package esquery

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

type mapTest struct {
	name string
	q    Mappable
	exp  map[string]interface{}
}

func runMapTests(t *testing.T, tests []mapTest) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := test.q.Map()

			// convert both maps to JSON in order to compare them. we do not
			// use reflect.DeepEqual on the maps as this doesn't always work
			exp, got, ok := sameJSON(test.exp, m)
			if !ok {
				t.Errorf("expected %s, got %s", exp, got)
			}
		})
	}
}

func sameJSON(a, b map[string]interface{}) (aJSON, bJSON []byte, ok bool) {
	aJSON, aErr := json.Marshal(a)
	bJSON, bErr := json.Marshal(b)

	if aErr != nil || bErr != nil {
		return aJSON, bJSON, false
	}

	ok = reflect.DeepEqual(aJSON, bJSON)
	return aJSON, bJSON, ok
}

type jsonTest struct {
	name    string
	q       json.Marshaler
	expJSON string
	expErr  error
}

func runJSONTests(t *testing.T, tests []jsonTest) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := test.q.MarshalJSON()
			assert.Equal(t, test.expErr, err)
			assert.Equal(t, test.expJSON, string(b))
		})
	}
}
