/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0 

	var height func(*TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		if left + right > diameter {
			diameter = left + right 
		}
		return 1 + max(left, right)
	}

	height(root)

	return diameter
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
