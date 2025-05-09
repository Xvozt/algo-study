package remove_duplicates_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_remove_duplications(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 2},
			result: 2,
		},
		{
			name:   "test 2",
			nums:   []int{1, 1, 2, 3, 4, 4, 5},
			result: 5,
		},
		{
			name:   "test 3",
			nums:   []int{1, 2, 3, 4, 5},
			result: 5,
		},
		{
			name:   "test 4",
			nums:   []int{},
			result: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)

			assert.Equal(t, tt.result, got)
		})
	}
}
