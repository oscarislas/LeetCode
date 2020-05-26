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

import "fmt"

func firstMissingPositive(nums []int) int {
	max := len(nums)
	if max <= 0 || (max == 1 && nums[0] != 1) {
		return 1
	}
	if max == 1 && nums[0] == 1 {
		return 2
	}

	for i, val := range nums {
		if val > max || val <= 0 {
			nums[i] = 1
			continue
		}
		if nums[val-1] > max || nums[val-1] <= 0 {
			nums[0] = nums[val-1]
		}
		nums[val-1] = -1
	}
	fmt.Println(nums)
	for i, val := range nums {
		if val > 0 {
			return i + 1
		}
	}
	if nums[0] > 0 {
		return nums[0]
	}
	return max
}
