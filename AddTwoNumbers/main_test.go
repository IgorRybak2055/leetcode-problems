package main

import (
	"testing"
)

type TestCase struct {
	list1        *ListNode
	list2        *ListNode
	expectedList *ListNode
}

func TestSolutionAddTwoNumber(t *testing.T) {
	testCases := []TestCase{
		{
			list1:
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			list2:
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expectedList:
			&ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			list1:
			&ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
			list2:
			&ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
			expectedList:
			&ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			list1:
			&ListNode{
				Val: 1,
				Next: nil,
			},
			list2:
			&ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: nil,
				},
			},
			expectedList:
			&ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			list1:
			&ListNode{
				Val: 4,
				Next: nil,
			},
			list2:
			&ListNode{
				Val: 5,
				Next: nil,
			},
			expectedList:
			&ListNode{
				Val: 9,
				Next: nil,
			},
		},
	}

	for idx, item := range testCases {
		answer := addTwoNumbers(item.list1, item.list2)

		for answer.Next != nil {
			if answer.Val != item.expectedList.Val {
				t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, answer.Val,
					item.expectedList.Val)
			}
			answer, item.expectedList = answer.Next, item.expectedList.Next
		}
	}
}
