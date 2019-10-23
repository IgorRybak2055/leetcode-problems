package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	head     *ListNode
	n        int
	expected *ListNode
}

func TestSolutionRemoveNthFromEnd(t *testing.T) {
	testCases := []TestCase{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			n: 2,
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  88,
					Next: nil,
				},
			},
			n: 4,
			expected: &ListNode{
				Val:  88,
				Next: nil,
			},
		},
		{
			head: &ListNode{
				Val:  5,
				Next: nil,
			},
			n:        4,
			expected: nil,
		},
	}

	for idx, item := range testCases {
		ans := removeNthFromEnd(item.head, item.n)

		if !reflect.DeepEqual(ans, item.expected) {
			t.Errorf("[ case %v ] : results not match\nGot:\n%#v\nExpected:\n%#v", idx, ans, item.expected)
		}
	}
}
