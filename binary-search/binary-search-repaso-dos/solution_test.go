package is_bad_version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_first_bad_version(t *testing.T) {
	tests := []struct {
		name   string
		chars  []byte
		target byte
		result byte
	}{
		{
			name:   "test 1",
			chars:  []byte{'c', 'f', 'j'},
			target: 'a',
			result: 'c',
		},
		{
			name:   "test 2",
			chars:  []byte{'c', 'f', 'j'},
			target: 'c',
			result: 'f',
		},
		{
			name:   "test 3",
			chars:  []byte{'c', 'f', 'j'},
			target: 'j',
			result: 'c',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestLetterGreaterThanTarget(tt.chars, tt.target)
			assert.Equal(t, tt.result, got)
		})
	}
}
