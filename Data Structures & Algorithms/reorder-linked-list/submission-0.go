/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
		return
	}

	slow , fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil 

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first := head
	second := prev

	for second != nil {
		f1 := first.Next
		f2 := second.Next

		first.Next = second
		second.Next = f1

		first = f1
		second = f2
	}
}

