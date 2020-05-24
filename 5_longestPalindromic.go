//https://leetcode.com/problems/longest-palindromic-substring/
//5. Longest Palindromic Substring

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// Example 1:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.

package main

func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}
	longestP := string(s[0])
	for index := range s {
		currentP := expand(s, index, index)
		if len(currentP) > len(longestP) {
			longestP = currentP
		}
		currentP = expand(s, index, index+1)
		if len(currentP) > len(longestP) {
			longestP = currentP
		}
	}
	return longestP
}

func expand(s string, x int, y int) string {
	for x >= 0 && y < len(s) && s[x] == s[y] {
		x--
		y++
	}
	return s[x+1 : y]
}
