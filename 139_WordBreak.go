//https://leetcode.com/problems/word-break/
//139. Word Break

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// determine if s can be segmented into a space-separated sequence of one or more dictionary words.
// Note:
//     The same word in the dictionary may be reused multiple times in the segmentation.
//     You may assume the dictionary does not contain duplicate words.
// Example 1:
// Input: s = "leetcode", wordDict = ["leet", "code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

package main

func wordBreak(s string, wordDict []string) bool {
	if len(s) <= 0 {
		return false
	}
	wordMap := make(map[string]struct{})
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	for i := range s {
		for j := 0; j <= i; j++ {
			if _, ok := wordMap[s[j:i+1]]; ok && dp[j] {
				dp[i+1] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
