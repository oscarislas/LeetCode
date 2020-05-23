//https://leetcode.com/problems/merge-k-sorted-lists/
//23. Merge k Sorted Lists

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
// Example:
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6
package main

//O(N log k)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}
	max := len(lists)
	interval := 1
	for interval < max {
		for i := 0; i < max-interval; i = i + (interval * 2) {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

//Recursive O(k*N)
func mergeKListsR(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	minIndex := minVal(lists)
	if minIndex == -1 {
		return nil
	}
	head := lists[minIndex]
	if head.Next != nil {
		lists[minIndex] = head.Next
	} else {
		lists[minIndex] = lists[len(lists)-1]
		lists = lists[:len(lists)-1]
	}
	head.Next = mergeKLists(lists)
	return head
}

func minVal(lists []*ListNode) int {
	var min *ListNode
	var minIndex = -1
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil && (min == nil || (*(lists[i])).Val < min.Val) {
			min = lists[i]
			minIndex = i
		}
	}
	return minIndex
}
