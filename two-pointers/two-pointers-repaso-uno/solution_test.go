package move_zeroes_in_place

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_remove_duplications(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result []int
	}{
		{
			name:   "test 1",
			nums:   []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
		{
			name:   "test 2",
			nums:   []int{1, 2, 3},
			result: []int{1, 2, 3},
		},
		{
			name:   "test 3",
			nums:   []int{0, 0, 0},
			result: []int{0, 0, 0},
		},
		{
			name:   "test 4",
			nums:   []int{0, 0, 1},
			result: []int{1, 0, 0},
		},
		{
			name:   "test 4",
			nums:   []int{1, 2, 0, 0},
			result: []int{1, 2, 0, 0},
		},
		{
			name:   "test 5",
			nums:   []int{},
			result: []int{},
		},
		{
			name:   "test 5",
			nums:   []int{0},
			result: []int{0},
		},
		{
			name:   "test 5",
			nums:   []int{5},
			result: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZeroes(tt.nums)
			assert.Equal(t, tt.result, got)
		})
	}
}
