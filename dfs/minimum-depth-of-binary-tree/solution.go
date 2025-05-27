package minimumdepthofbinarytree

/*

Дано бинарное дерево.
Верни минимальную глубину — это количество узлов от корня до ближайшего листа (то есть узла без детей).

Input:
        1
       /
      2
Output: 2  // путь: 1 → 2 (но 2 — не лист, поэтому глубина — больше)

Input:
        3
       / \
      9  20
         / \
        15  7

Output: 2
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	if node.Left == nil {
		return 1 + minimumDepth(node.Right)
	}
	if node.Right == nil {
		return 1 + minimumDepth(node.Left)
	}

	return 1 + min(minimumDepth(node.Left), minimumDepth(node.Right))

}
