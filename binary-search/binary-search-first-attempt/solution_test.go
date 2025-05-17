package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search_insert_index(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			result: 2,
		},
		{
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			result: 1,
		},
		{
			name:   "test 3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			result: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsertPosition(tt.nums, tt.target)
			assert.Equal(t, tt.result, got)
		})
	}
}
