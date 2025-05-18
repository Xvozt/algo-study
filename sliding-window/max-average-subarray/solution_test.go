package max_average_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longest_subarray_after_deletion(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		result float64
	}{
		{
			name:   "test 1",
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			result: 12.75,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAverageSubarray(tt.nums, tt.k)
			assert.Equal(t, tt.result, got)
		})
	}
}
