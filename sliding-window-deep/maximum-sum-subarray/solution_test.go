package maximum_sum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximum_sum_subarray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{2, 1, 5, 1, 3, 2},
			k:      3,
			result: 9,
		},
		{
			name:   "test 2",
			nums:   []int{},
			k:      3,
			result: 0,
		},
		{
			name:   "test 3",
			nums:   []int{1},
			k:      3,
			result: 1,
		},
		{
			name:   "test 4",
			nums:   []int{1, 2},
			k:      3,
			result: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSum(tt.nums, tt.k)
			assert.Equal(t, tt.result, got)
		})
	}
}
