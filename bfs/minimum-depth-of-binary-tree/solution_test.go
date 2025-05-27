package minimumdepthofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_min_depth_binary_tree(t *testing.T) {
	tests := []struct {
		name   string
		tree   *TreeNode
		result int
	}{
		{
			name: "test 1",
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			result: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDepth(tt.tree)
			assert.Equal(t, tt.result, got)
		})
	}
}
