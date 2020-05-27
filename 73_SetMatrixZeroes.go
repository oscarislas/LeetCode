//https://leetcode.com/problems/set-matrix-zeroes/
//73. Set Matrix Zeroes

// Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
// Example 1:
// Input:
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// Output:
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

package main

func setZeroes(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}

	firstRowHasZero := false

	for row := 0; row < len(matrix); row++ {
		firstRowHasZero = firstRowHasZero || matrix[row][0] == 0
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				matrix[row][0] = 0
			}
		}
	}
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[0][col] = 0
		}
	}
	if firstRowHasZero {
		for row := 0; row < len(matrix); row++ {
			matrix[row][0] = 0
		}
	}
}
