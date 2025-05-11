package max_consecutive_ones_iii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_count_max_consecutive_ones_iii(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:      2,
			result: 6,
		},
		{
			name:   "test 2",
			nums:   []int{0, 0, 0, 0, 0, 0},
			k:      2,
			result: 2,
		},
		{
			name:   "test 3",
			nums:   []int{1, 1, 1, 1, 1, 1},
			k:      2,
			result: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxConsecutiveOnesIII(tt.nums, tt.k)
			assert.Equal(t, tt.result, got)
		})
	}
}
