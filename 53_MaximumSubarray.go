//https://leetcode.com/problems/maximum-subarray/
//53. Maximum Subarray

// Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and return its sum.
// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:
// If you have figured out the O(n) solution,
// try coding another solution using the divide and conquer approach,
// which is more subtle.

package main

func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for _, val := range nums[1:] {
		currSum = max(val, currSum+val)
		maxSum = max(currSum, maxSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
