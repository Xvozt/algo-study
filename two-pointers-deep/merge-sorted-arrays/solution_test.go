package mergesortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse_string(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			name:   "test 1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.nums1, tt.m, tt.nums2, tt.n)

			assert.Equal(t, tt.result, got)
		})
	}
}
