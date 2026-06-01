/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	copyMap := make(map[*Node]*Node)

	curr := head

	for curr != nil {
		copyMap[curr] = &Node{
			Val : curr.Val,
		}
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		copyMap[curr].Next = copyMap[curr.Next]
		copyMap[curr].Random = copyMap[curr.Random]
		curr = curr.Next
	}

	return copyMap[head]
}
