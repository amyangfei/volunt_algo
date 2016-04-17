package string

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var pre *ListNode = nil
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	slow.Next = pre

	for slow.Next != nil {
		slow = slow.Next
		if slow.Val != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}
