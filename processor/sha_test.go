package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToHexSHA(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		result string
	}{
		{
			name:   "value 0",
			input:  0,
			result: "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
		},
		{
			name:   "value 10",
			input:  10,
			result: "4a44dc15364204a80fe80e9039455cc1608281820fe2b24f1e5233ade6af1dd5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToHexSHA(tt.input)
			assert.Equal(t, tt.result, result)
		})
	}
}
