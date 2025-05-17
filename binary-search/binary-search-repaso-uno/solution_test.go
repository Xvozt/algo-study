package is_bad_version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_first_bad_version(t *testing.T) {
	tests := []struct {
		name   string
		num    int
		bad    int
		result int
	}{
		{
			name:   "test 1",
			num:    5,
			bad:    4,
			result: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstBadVersion(tt.num)
			assert.Equal(t, tt.result, got)
		})
	}
}
