/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}

	q := []*TreeNode{root}
	result := []int{}
	for len(q) > 0 {
		levelSize := len(q)
		
		for i := 0 ; i < levelSize; i++ {
			node := q[0]
			q = q[1:]

			if i == levelSize - 1{
				result = append(result,node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return result
}
