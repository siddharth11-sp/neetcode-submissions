/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var validate func(*TreeNode, int, int) bool 

	validate = func(root *TreeNode, minVal, maxVal int) bool {
		if root == nil {
			return true
		}

		if root.Val <= minVal || root.Val >= maxVal {
			return false
		}

		return validate(root.Left, minVal, root.Val) &&
		validate(root.Right, root.Val, maxVal)
	}

	return validate(root, math.MinInt, math.MaxInt )
}
