//https://leetcode.com/problems/spiral-matrix/
//54. Spiral Matrix

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
// Example 1:
// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]

package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 0 {
		return nil
	}
	ans := make([]int, 0)
	n := len(matrix)
	m := len(matrix[0])
	row1, col1, row2, col2 := 0, 0, n-1, m-1
	for row1 <= row2 && col1 <= col2 {
		for col := col1; col <= col2; col++ {
			ans = append(ans, matrix[row1][col])
		}
		for row := row1 + 1; row <= row2; row++ {
			ans = append(ans, matrix[row][col2])
		}
		if row1 < row2 && col1 < col2 {
			for col := col2 - 1; col >= col1; col-- {
				ans = append(ans, matrix[row2][col])
			}
			for row := row2 - 1; row > row1; row-- {
				ans = append(ans, matrix[row][col1])
			}
		}
		col1++
		col2--
		row1++
		row2--
	}
	return ans
}
