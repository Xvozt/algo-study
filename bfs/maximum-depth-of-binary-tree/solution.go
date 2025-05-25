package maximumdepthofbinarytree

/*
Дано бинарное дерево.
Верни максимальную глубину дерева.

Глубина — количество узлов по самому длинному пути от корня до листа.

Input:
        3
       / \
      9  20
         / \
        15  7

Output: 3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}
