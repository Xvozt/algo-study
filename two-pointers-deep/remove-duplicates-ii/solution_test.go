package removeduplicatesii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse_string(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 1, 2, 2, 3},
			result: 5,
		},

		{
			name:   "test 2",
			nums:   []int{1, 2, 3, 4},
			result: 4,
		},
		{
			name:   "test 3",
			nums:   []int{},
			result: 0,
		},
		{
			name:   "test 4",
			nums:   []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			result: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)

			assert.Equal(t, tt.result, got)
		})
	}
}
