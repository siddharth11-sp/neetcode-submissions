/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    inorderMap := make(map[int]int, len(inorder))

	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(int , int) *TreeNode 
	preIdx := 0
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootVal := preorder[preIdx]
		preIdx++
		root := &TreeNode{
			Val : rootVal,
		}

		mid := inorderMap[rootVal]

		root.Left = build(left, mid -1)
		root.Right = build(mid+1,right)

		return root

	}
	return build(0, len(inorder)-1)
}
