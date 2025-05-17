package longestsubstringwithoutrepeating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longest_substring_without_repeating(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		result int
	}{
		{
			name:   "test 1",
			str:    "abcabcbb",
			result: 3,
		},
		{
			name:   "test 2",
			str:    "bbbbb",
			result: 1,
		},
		{
			name:   "test 3",
			str:    "pwwkew",
			result: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubstringWithoutRepeating(tt.str)
			assert.Equal(t, tt.result, got)
		})
	}
}
