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

func minimumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}
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
