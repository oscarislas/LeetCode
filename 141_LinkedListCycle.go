//https://leetcode.com/problems/linked-list-cycle/
//141. Linked List Cycle

// Given a linked list, determine if it has a cycle in it.
// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed)
// in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the second node.

package main

func hasCycle(head *ListNode) bool {
	for hare, tortoise := head, head; hare != nil && hare.Next != nil; {
		hare, tortoise = hare.Next.Next, tortoise.Next
		if hare == tortoise {
			return true
		}
	}
	return false
}
