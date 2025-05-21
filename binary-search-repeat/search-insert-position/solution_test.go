package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search_insert_posittion(t *testing.T) {
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
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			result: 4,
		},
		{
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			result: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			assert.Equal(t, tt.result, got)
		})
	}
}
