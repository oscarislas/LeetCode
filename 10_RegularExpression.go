//https://leetcode.com/problems/regular-expression-matching/
//10. Regular Expression Matching

// Given an input string (s) and a pattern (p),
// implement regular expression matching with support for '.' and '*'.
// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).
// Note:
//     s could be empty and contains only lowercase letters a-z.
//     p could be empty and contains only lowercase letters a-z, and characters like . or *.
// Example 1:
// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa"

package main

func isMatch(s string, p string) bool {
	if len(p) <= 0 {
		return len(s) <= 0
	}
	if s == p {
		return true
	}
	firstMatch := len(s) > 0 && (s[0] == p[0] || (len(s) > 0 && p[0] == '.'))
	if len(p) > 1 && p[1] == '*' {
		if firstMatch {
			return isMatch(s, p[2:]) || isMatch(s[1:], p)
		}
		return isMatch(s, p[2:])
	}
	if firstMatch {
		return true && isMatch(s[1:], p[1:])
	}
	return false
}
