//https://leetcode.com/problems/binary-tree-level-order-traversal/
//102. Binary Tree Level Order Traversal

// Given a binary tree,
// return the level order traversal of its nodes' values.
// (ie, from left to right, level by level).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		level_length := len(queue)
		result = append(result, make([]int, 0))
		for i := 0; i < level_length; i++ {
			node := queue[0]
			queue = queue[1:]
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}
	return result
}
