package main

import "fmt"

func main() {
	var a = &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ans := removeNthFromEnd(a, 2)

	for ans.Next != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := remove(head, n)
	if cnt < 1 {
		return nil
	}
	if cnt < n {
		return head.Next
	}

	return head
}

func remove(node *ListNode, n int) int {
	if node.Next == nil {
		return 0
	}

	cnt := remove(node.Next, n)
	cnt++

	if cnt == n {
		node.Next = node.Next.Next
	}

	return cnt
}