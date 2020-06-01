//https://leetcode.com/problems/reorder-list/
//143. Reorder List

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
// You may not modify the values in the list's nodes, only nodes itself may be changed.
// Example 1:
// Given 1->2->3->4, reorder it to 1->4->2->3.

package main

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	first, second := head, head
	for second != nil && second.Next != nil {
		first, second = first.Next, second.Next.Next
	}

	var prev *ListNode
	for first != nil {
		first.Next, prev, first = prev, first, first.Next
	}

	for prev.Next != nil {
		head.Next, head = prev, head.Next
		prev.Next, prev = head, prev.Next
	}
}
