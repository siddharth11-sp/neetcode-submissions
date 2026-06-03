/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    
	// first need to check if the list contians the k nodes.
	curr := head
	for i := 0; i < k ; i++ {
		if curr == nil {
			return head
		}
		curr = curr.Next
	}

	var prev *ListNode
	curr = head

	for i := 0; i < k ; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr =next
	}

	head.Next = reverseKGroup(curr, k)
	return prev
}
