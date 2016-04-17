/*
1.5 单链表回文
*/
package string

import (
	"testing"
)

func Array2LinkedList(nums []int) *ListNode {
	dummy := &ListNode{0, nil}
	head := dummy
	for _, num := range nums {
		node := &ListNode{num, nil}
		head.Next = node
		head = node
	}
	return dummy.Next
}

func testCase(t *testing.T, nums []int, expected bool) {
	if IsPalindrome(Array2LinkedList(nums)) != expected {
		t.Errorf("err for %v", nums)
	}
}

func TestIsPalindrome(t *testing.T) {
	test1 := []int{1, 2, 1}
	testCase(t, test1, true)

	test2 := []int{1, 2, 3, 0, 0, 3, 2, 1}
	testCase(t, test2, true)

	test3 := []int{1, 2, 2}
	testCase(t, test3, false)

	test4 := []int{1, 1}
	testCase(t, test4, true)

	test5 := []int{1}
	testCase(t, test5, true)
}
