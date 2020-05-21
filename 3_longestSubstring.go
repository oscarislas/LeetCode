//https://leetcode.com/problems/longest-substring-without-repeating-characters/
// 3. Longest Substring Without Repeating Characters

// Given a string, find the length of the longest substring without repeating characters.
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
package main

func lengthOfLongestSubstring(s string) int {
	//empty string
	if len(s) <= 0 {
		return 0
	}
	//init values
	max, count, x, y := 1, 1, 0, 0
	existingRunes := make(map[byte]int)
	existingRunes[s[y]]++

	//logic
	for {
		if existingRunes[s[y]] > 1 {
			existingRunes[s[x]]--
			x++
			count--
			continue
		}
		if max < count {
			max = count
		}
		count++
		y++
		if y >= len(s) {
			break
		}
		existingRunes[s[y]]++
	}
	return max
}
