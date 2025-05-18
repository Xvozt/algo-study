package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse_string(t *testing.T) {
	tests := []struct {
		name   string
		str    []byte
		result []byte
	}{
		{
			name:   "test 1",
			str:    []byte{'h', 'e', 'l', 'l', 'o'},
			result: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.str)

			assert.Equal(t, tt.result, got)
		})
	}
}
