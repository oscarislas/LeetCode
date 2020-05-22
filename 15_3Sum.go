//https://leetcode.com/problems/3sum/
//15. 3Sum

// Given an array nums of n integers, are there elements a, b, c
// in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
// Note:
// The solution set must not contain duplicate triplets.
// Example:
// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	response := [][]int{}
	if len(nums) < 3 {
		return response
	}
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -n
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[j] + nums[k]
			if sum == target {
				response = append(response, []int{nums[i], nums[j], nums[k]})
			} else if sum < target {
				for {
					j++
					if j > k || nums[j] != nums[j-1] {
						break
					}
				}
				continue
			}
			for {
				k--
				if k < j || nums[k] != nums[k+1] {
					break
				}
			}

		}
	}
	return response
}
