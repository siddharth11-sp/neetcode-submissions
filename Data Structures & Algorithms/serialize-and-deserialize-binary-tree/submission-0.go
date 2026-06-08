/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    result := []string{}

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result = append(result, "N")
			return
		}

		result = append(result, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
		return 
	}
	dfs(root)
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    values := strings.Split(data, ",")

	idx := 0 

	var dfs func() *TreeNode

	dfs = func() *TreeNode {
		if values[idx] == "N" {
			idx++
			return nil
		}

		val, _ := strconv.Atoi(values[idx])
		idx++

		root := &TreeNode{
			Val : val,
		}

		root.Left = dfs()
		root.Right = dfs()

		return root
	}
	return dfs() 
}
