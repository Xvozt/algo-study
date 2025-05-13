package longest_subarray_after_deletion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longest_subarray_after_deletion(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 0, 1},
			result: 3,
		},
		{
			name:   "test 2",
			nums:   []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			result: 5,
		},
		{
			name:   "test 3",
			nums:   []int{0, 0, 0, 0, 0},
			result: 0,
		},
		{
			name:   "test 4",
			nums:   []int{1, 1, 1},
			result: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubArrayAfterDeletion(tt.nums)
			assert.Equal(t, tt.result, got)
		})
	}
}
