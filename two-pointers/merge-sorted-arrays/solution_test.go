package mergesortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_remove_duplications(t *testing.T) {
	tests := []struct {
		name    string
		numsOne []int
		m       int
		numsTwo []int
		n       int
		result  []int
	}{
		{
			name:    "test 1",
			numsOne: []int{1, 1, 2, 0, 0, 0},
			m:       3,
			numsTwo: []int{2, 5, 6},
			n:       3,
			result:  []int{1, 1, 2, 2, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeSortedArrays(tt.numsOne, tt.m, tt.numsTwo, tt.n)

			assert.Equal(t, tt.result, got)
		})
	}
}
