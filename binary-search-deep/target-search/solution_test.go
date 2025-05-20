package targetsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_target_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			result: 4,
		},
		{
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			result: -1,
		},
		{
			name:   "test 3",
			nums:   []int{-5, 1, 3, 5, 6},
			target: -1,
			result: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := targetSearch(tt.nums, tt.target)
			assert.Equal(t, tt.result, got)
		})
	}
}
