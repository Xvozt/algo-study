package sorted_squares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sorted_squares(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result []int
	}{
		{
			name:   "test 1",
			nums:   []int{-4, -1, 0, 3, 10},
			result: []int{0, 1, 9, 16, 100},
		},
		{
			name:   "test 2",
			nums:   []int{1, 2, 3, 4},
			result: []int{1, 4, 9, 16},
		},
		{
			name:   "test 3",
			nums:   []int{},
			result: []int{},
		},
		{
			name:   "test 4",
			nums:   []int{-1},
			result: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.nums)

			assert.Equal(t, tt.result, got)
		})
	}
}
