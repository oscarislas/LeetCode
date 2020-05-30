//https://leetcode.com/problems/binary-tree-maximum-path-sum/
//124. Binary Tree Maximum Path Sum

// Given a non-empty binary tree, find the maximum path sum.
// For this problem, a path is defined as any sequence of nodes
// from some starting node to any node in the tree along the parent-child connections.
// The path must contain at least one node and does not need to go through the root.
// Example 1:
// Input: [1,2,3]
//        1
//       / \
//      2   3
// Output: 6
// Example 2:
// Input: [-10,9,20,null,null,15,7]
//    -10
//    / \
//   9  20
//     /  \
//    15   7
// Output: 42

package main

import "math"

func maxPathSum(root *TreeNode) int {
	max := math.MinInt64
	maxPathSumHelper(root, &max)
	return max
}

func maxPathSumHelper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	var lMax, rMax int
	if lMax = maxPathSumHelper(root.Left, max); lMax < 0 {
		lMax = 0
	}
	if rMax = maxPathSumHelper(root.Right, max); rMax < 0 {
		rMax = 0
	}
	newPath := root.Val + lMax + rMax
	if *max < newPath {
		*max = newPath
	}

	if lMax > rMax {
		return root.Val + lMax
	}
	return root.Val + rMax
}
