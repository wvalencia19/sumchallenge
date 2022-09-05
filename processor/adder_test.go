package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPayloadNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		total int
	}{
		{
			name:  "array",
			input: []byte(`{"data": [1,2,3,4]}`),
			total: 10,
		},
		{
			name:  "map",
			input: []byte(`{"data": {"a":6,"b":4}}`),
			total: 10,
		},
		{
			name:  "nested array",
			input: []byte(`{"data": [[[2]]]}`),
			total: 2,
		},
		{
			name:  "nested map",
			input: []byte(`{"data": {"a":{"b":4},"c":-2}}`),
			total: 2,
		},
		{
			name:  "different structures",
			input: []byte(`{"data": {"a":{"b":4},"c":-2}, "data1": [[[2]]], "data2": {"a":6,"b":4}, "data3": [-1,{"a":1, "b":"light"}]}`),
			total: 14,
		},
		{
			name:  "without numbers",
			input: []byte(`{"data": "this is a test"}`),
			total: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapInterface, err := UnmarshallToInterface(tt.input)
			assert.NoError(t, err)
			total := AddPayloadNumbers(mapInterface)
			assert.Equal(t, tt.total, total)
		})
	}
}
