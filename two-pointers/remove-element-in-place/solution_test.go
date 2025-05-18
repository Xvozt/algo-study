package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_remove_duplications(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		val    int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{3, 2, 2, 3},
			val:    3,
			result: 2,
		},
		{
			name:   "test 2",
			nums:   []int{2, 2, 2, 3},
			val:    2,
			result: 1,
		},
		{
			name:   "test 3",
			nums:   []int{},
			val:    1,
			result: 0,
		},
		{
			name:   "test 4",
			nums:   []int{1, 1, 1, 1},
			val:    1,
			result: 0,
		},
		{
			name:   "test 4",
			nums:   []int{1, 2, 3, 4},
			val:    5,
			result: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			assert.Equal(t, tt.result, got)
		})
	}
}
