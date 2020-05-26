//https://leetcode.com/problems/insert-interval/
//57. Insert Interval

// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
// You may assume that the intervals were initially sorted according to their start times.
// Example 1:
// Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
// Output: [[1,5],[6,9]]

package main

func insert(intervals [][]int, newInterval []int) [][]int {
	merged := make([][]int, 0)
	i, n := 0, len(intervals)
	for i := range intervals {
		if intervals[i][0] > newInterval[0] {
			break
		}
		merged = append(merged, intervals[i])
	}

	if len(merged) == 0 || merged[len(merged)-1][1] < newInterval[0] {
		merged = append(merged, newInterval)
	} else {
		lastMerged := merged[len(merged)-1]
		if lastMerged[1] < newInterval[1] {
			lastMerged[1] = newInterval[1]
		}
	}

	for ; i < n; i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		}
		lastMerged := merged[len(merged)-1]
		if lastMerged[1] < intervals[i][1] {
			lastMerged[1] = intervals[i][1]
		}
	}

	return merged
}
