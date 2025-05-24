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

func maximumDepthOfBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(maximumDepthOfBinaryTree(node.Left), maximumDepthOfBinaryTree(node.Right))
}
