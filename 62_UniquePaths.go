//https://leetcode.com/problems/unique-paths/
//62. Unique Paths

// A robot is located at the top-left corner of a m x n grid
//(marked 'R' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid
//(marked '*' in the diagram below).
// How many possible unique paths are there?
// [R][ ][ ][ ][ ][ ]
// [ ][ ][ ][ ][ ][ ]
// [ ][ ][ ][ ][ ][*]
// Example 1:
// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Right -> Down
// 2. Right -> Down -> Right
// 3. Down -> Right -> Right
// Example 2:
// Input: m = 7, n = 3
// Output: 28

package main

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	valueMap := make([][]int, m)
	for i := range valueMap {
		valueMap[i] = make([]int, n)
		valueMap[i][0] = 1
	}
	for j := range valueMap[0] {
		valueMap[0][j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			valueMap[i][j] = valueMap[i-1][j] + valueMap[i][j+1]
		}
	}

	return valueMap[n-1][m-1]
}
