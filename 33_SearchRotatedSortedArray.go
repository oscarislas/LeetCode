//https://leetcode.com/problems/search-in-rotated-sorted-array/
//33. Search in Rotated Sorted Array

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm's runtime complexity must be in the order of O(log n).
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:
// Input: nums = [7,8,0,1,2,3], target = 3
// Output: -1

package main

func search(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		pivot := left + ((right - left) / 2)
		if nums[pivot] == target {
			return pivot
		}
		if nums[left] <= nums[pivot] {
			//balanced
			if target >= nums[left] && target < nums[pivot] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		} else {
			//not balanced
			if target > nums[pivot] && target <= nums[length-1] {
				left = pivot + 1
			} else {
				right = pivot - 1
			}
		}
	}
	return -1
}

func searchDoublePass(nums []int, target int) int {
	length := len(nums) - 1
	if length < 0 {
		return -1
	}
	if nums[0] <= nums[length] {
		return binarySearch(0, length, nums, target)
	}
	pivot := searchPivot(nums)
	if target == nums[pivot] {
		return pivot
	}
	if target >= nums[0] {
		return binarySearch(0, pivot, nums, target)
	}
	return binarySearch(pivot, length, nums, target)
}

func searchPivot(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := left + ((right - left) / 2)
		if nums[mid] > nums[right] {
			left = mid + 1
			continue
		}
		right = mid
	}
	return left
}

func binarySearch(start int, end int, nums []int, target int) int {
	left, right := start, end
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return -1
}
