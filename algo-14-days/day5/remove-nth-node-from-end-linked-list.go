package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy *ListNode = &ListNode{0, head}
	var slow, fast *ListNode = dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	var res *ListNode = removeNthFromEnd(head, 2)
	fmt.Println(res)
}
