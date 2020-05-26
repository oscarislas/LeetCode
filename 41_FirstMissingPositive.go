//https://leetcode.com/problems/first-missing-positive/solution/
//41. First Missing Positive

// Given an unsorted integer array, find the smallest missing positive integer.
// Example 1:
// Input: [1,2,0]
// Output: 3

// Example 2:
// Input: [3,4,-1,1]
// Output: 2

package main

func firstMissingPositive(nums []int) int {
	max := len(nums)
	if max <= 0 {
		return 1
	}
	for i := 0; i < max; {
		j := nums[i] - 1
		if j >= 0 && j < max && j != nums[j]-1 {
			nums[i], nums[j] = nums[j], nums[i]
			continue
		}
		i++
	}
	for i, val := range nums {
		if val != i+1 {
			return i + 1
		}
	}
	return max + 1
}
