/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt

	var dfs func(*TreeNode) int 

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(dfs(root.Left),0)
		right := max(dfs(root.Right), 0)

		currMax := root.Val + left + right 

		if currMax > maxSum {
			maxSum = currMax
		}

		return root.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}
