//https://leetcode.com/problems/word-search/
//79. Word Search

// Given a 2D board and a word, find if the word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
// Example:
// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// Given word = "ABCCED", return true.
// Given word = "SEE", return true.
// Given word = "ABCB", return false.

package main

func exist(board [][]byte, word string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == word[0] {
				temp := board[row][col]
				board[row][col] = byte('#')
				if isCompleteWord(board, word[1:], row, col) {
					return true
				}
				board[row][col] = temp
			}
		}
	}
	return false
}

func isCompleteWord(board [][]byte, word string, row, col int) bool {
	if len(word) <= 0 {
		return true
	}
	rows, cols := []int{1, 0, -1, 0}, []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		r := row + rows[i]
		c := col + cols[i]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && word[0] == board[r][c] {
			temp := board[r][c]
			board[r][c] = byte('#')
			if isCompleteWord(board, word[1:], r, c) {
				return true
			}
			board[r][c] = temp
		}
	}
	return false
}
