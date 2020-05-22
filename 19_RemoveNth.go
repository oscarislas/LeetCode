//https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//19. Remove Nth Node From End of List

// Given a linked list, remove the n-th node from the end of list and return its head.
// Example:
// Given linked list: 1->2->3->4->5, and n = 2.
// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:
// Given n will always be valid.
// Follow up:
// Could you do this in one pass?

package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	temp := &ListNode{
		Next: head,
	}
	x, y := temp, temp
	for i := 1; i <= n; i++ {
		y = y.Next
	}
	for y.Next != nil {
		x = x.Next
		y = y.Next
	}
	x.Next = x.Next.Next
	return temp.Next
}

//O(n) with O(n) space
func removeNthFromEnd_slice(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	s := []*ListNode{}
	for {
		s = append(s, head)
		head = head.Next
		if head == nil {
			break
		}
	}
	x := len(s) - n
	if x > 0 {
		s[x-1].Next = s[x].Next
		return s[0]
	}
	return s[1]
}
