//https://leetcode.com/problems/longest-consecutive-sequence/
//128. Longest Consecutive Sequence

// Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
// Your algorithm should run in O(n) complexity.
// Example:
// Input: [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

package main

func longestConsecutive(nums []int) int {
	hash := make(map[int]struct{})
	max := 0
	for _, num := range nums {
		hash[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := hash[num-1]; !ok {
			n := num
			streak := 1
			for {
				if _, ok2 := hash[n+1]; !ok2 {
					break
				}
				n++
				streak++
			}
			if max < streak {
				max = streak
			}
		}
	}
	return max
}
