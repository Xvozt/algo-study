package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_count_max_consecutive_ones(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 0, 1, 1, 1},
			result: 3,
		},
		{
			name:   "test 2",
			nums:   []int{0, 0, 0, 0, 0},
			result: 0,
		},
		{
			name:   "test 3",
			nums:   []int{1, 1, 1, 0, 1, 1},
			result: 3,
		},
		{
			name:   "test 4",
			nums:   []int{},
			result: 0,
		},
		{
			name:   "test 5",
			nums:   []int{1},
			result: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countMaxConsecutiveOnes(tt.nums)
			assert.Equal(t, tt.result, got)
		})
	}
}
