//https://leetcode.com/problems/jump-game/
//55. Jump Game

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Determine if you are able to reach the last index.
// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

package main

func canJump(nums []int) bool {
	last := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last == 0
}

func canJumpR(nums []int) bool {
	if len(nums) <= 1 || nums[0] >= len(nums) {
		return true
	}
	for i := nums[0]; i > 1; i-- {
		if canJumpR(nums[i:]) {
			return true
		}
	}
	return false
}
