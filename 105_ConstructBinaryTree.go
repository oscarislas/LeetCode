//https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//105. Construct Binary Tree from Preorder and Inorder Traversal

// Given preorder and inorder traversal of a tree, construct the binary tree.
// Note:
// You may assume that duplicates do not exist in the tree.
// For example, given
// preorder = [3,9,20,15,7]
// inorder = [9,3,15,20,7]
// Return the following binary tree:
//     3
//    / \
//   9  20
//     /  \
//    15   7

package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rVal := preorder[0]
	rIx := 0
	for i, val := range inorder {
		if val == rVal {
			rIx = i
		}
	}
	root := &TreeNode{Val: rVal}
	root.Left = buildTree(preorder[1:rIx+1], inorder[:rIx])
	root.Right = buildTree(preorder[rIx+1:], inorder[rIx+1:])

	return root
}
