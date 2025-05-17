package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valid_palindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "test 1",
			input:  "A man, a plan, a canal: Panama",
			result: true,
		},
		{
			name:   "test 2",
			input:  "race a car",
			result: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validPalindrome(tt.input)
			assert.Equal(t, tt.result, got)
		})
	}
}
