package firstandlastposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_first_and_last_position(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		result []int
	}{
		{
			name:   "test 1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			result: []int{3, 4},
		},
		{
			name:   "test 2",
			nums:   []int{15, 7, 7, 8, 8, 10},
			target: 6,
			result: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			assert.Equal(t, tt.result, got)
		})
	}
}
