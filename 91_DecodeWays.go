//https://leetcode.com/problems/decode-ways/
//91. Decode Ways

// A message containing letters from A-Z is being encoded to numbers
// using the following mapping:
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits,
// determine the total number of ways to decode it.
// Example 1:
// Input: "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).

package main

import "strconv"

func numDecodings(s string) int {
	if len(s) <= 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	i, j := 1, 1
	total := 1
	for x := 1; x < len(s); x++ {
		total = 0
		if s[x] != '0' {
			total += i
		}
		if n, err := strconv.Atoi(s[x-1 : x+1]); err == nil && n >= 10 && n <= 26 {
			total += j
		}
		i, j = total, i
	}
	return total
}
