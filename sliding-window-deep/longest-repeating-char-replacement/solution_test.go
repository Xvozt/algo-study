package longestrepeatingcharreplacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longest_subarray_after_deletion(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		k      int
		result int
	}{
		{
			name:   "test 1",
			str:    "ABAB",
			k:      2,
			result: 4,
		},
		{
			name:   "test 1",
			str:    "AABABBA",
			k:      1,
			result: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestRepeatingCharReplacement(tt.str, tt.k)
			assert.Equal(t, tt.result, got)
		})
	}
}
