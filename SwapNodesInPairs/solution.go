package main

import (
	"fmt"
)


type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	t := &ListNode{
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

	a := swapPairs(t)

	for {
		fmt.Println(a.Val)

		if a.Next == nil {
			break
		}
		a = a.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	result := head.Next
	head.Next = swapPairs(head.Next.Next)
	result.Next = head

	return result
}