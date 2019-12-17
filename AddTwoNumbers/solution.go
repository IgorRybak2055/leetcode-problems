package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main()  {
	// for travis CI
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return do(l1, l2, 0)
}

func do(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	out := &ListNode{}
	var val int
	if l1.Val+l2.Val > 9 {
		val = (l1.Val + l2.Val) %10 + carry
		carry = 1
	} else {
		val = l1.Val + l2.Val
		if val == 9 && carry == 1 {
			val = 0
			carry = 1
		} else if val < 9 && carry == 1 {
			val += carry
			carry = 0
		}

	}
	out.Val = val
	if l1.Next != nil && l2.Next != nil {
		out.Next = do(l1.Next, l2.Next, carry)
	}
	if l1.Next != nil && l2.Next == nil {
		out.Next = do(l1.Next, &ListNode{Val:0}, carry)
	}
	if l1.Next == nil && l2.Next != nil {
		out.Next = do(&ListNode{Val:0}, l2.Next, carry)
	}
	if l1.Next == nil && l2.Next == nil{
		if carry == 1 {
			out.Next = &ListNode{Val: carry}
		}
	}
	return out
}
