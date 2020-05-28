//https://leetcode.com/problems/merge-two-sorted-lists/
//21. Merge Two Sorted Lists

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
// Example:
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

//Definition for singly-linked list.

package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}

//non-recursive
func mergeTwoListsNR(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, last *ListNode = nil, nil
	if (*l2).Val >= (*l1).Val {
		head = l1
		last = l1
		l1 = head.Next
	} else {
		head = l2
		last = l2
		l2 = head.Next
	}
	for l1 != nil || l2 != nil {
		if l2 != nil && (l1 == nil || (*l2).Val < (*l1).Val) {
			last.Next = l2
			last = last.Next
			l2 = l2.Next
			continue
		}
		last.Next = l1
		last = last.Next
		l1 = l1.Next

	}
	return head
}

func mergeTwoListsHelper(x []int, y []int) []int {
	var l1 *ListNode
	var l2 *ListNode
	if len(x) > 0 {
		l1 = &ListNode{
			Val: x[0],
		}
		n := l1
		x = x[1:]
		for len(x) > 0 {
			n.Next = &ListNode{
				Val: x[0],
			}
			n = n.Next
			x = x[1:]
		}
	}

	if len(y) > 0 {
		l2 = &ListNode{
			Val: y[0],
		}
		n := l2
		y = y[1:]
		for len(y) > 0 {
			n.Next = &ListNode{
				Val: y[0],
			}
			n = n.Next
			y = y[1:]
		}
	}

	result := []int{}
	r := mergeTwoLists(l1, l2)
	for r != nil {
		result = append(result, r.Val)
		r = r.Next
	}
	return result
}
