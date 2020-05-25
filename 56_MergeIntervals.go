//https://leetcode.com/problems/merge-intervals/
//56. Merge Intervals

// Given a collection of intervals, merge all overlapping intervals.
// Example 1:
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

package main

import "sort"

func merge(intervals [][]int) [][]int {
	merged := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, interval := range intervals {
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
			continue
		}
		lastMerged := merged[len(merged)-1]
		if lastMerged[1] < interval[1] {
			lastMerged[1] = interval[1]
		}
	}
	return merged
}
