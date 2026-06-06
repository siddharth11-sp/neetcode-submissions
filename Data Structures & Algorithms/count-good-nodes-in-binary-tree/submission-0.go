/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var dfs func(*TreeNode, int) int

	dfs = func (root *TreeNode, maxSofar int) int {
		if root == nil {
			return 0
		}

		count := 0 

		if root.Val >= maxSofar {
			count = 1
			maxSofar = root.Val
		}

		count += dfs(root.Left, maxSofar)
		count += dfs(root.Right, maxSofar)

		return count
	}

	return dfs(root, root.Val)
}
