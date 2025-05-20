package minimuminrotatedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimum_in_rotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{3, 4, 5, 1, 2},
			result: 1,
		},
		{
			name:   "test 2",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			result: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumInRotated(tt.nums)
			assert.Equal(t, tt.result, got)
		})
	}
}
