//https://leetcode.com/problems/minimum-window-substring/
//76. Minimum Window Substring

// Given a string S and a string T, find the minimum window in S which will contain
// all the characters in T in complexity O(n).
// Example:
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"
// Note:
//     If there is no such window in S that covers all characters in T, return the empty string "".
// 	If there is such window, you are guaranteed that there will always be
// 	only one unique minimum window in S.

//AADCBBCC
//AABC
//AADCB

package main

func minWindow(s string, t string) string {
	if len(s) <= 0 || len(t) <= 0 || len(t) > len(s) {
		return ""
	}
	tHash, sHash := make(map[byte]int, len(t)), make(map[byte]int, len(s))
	for x := range t {
		tHash[t[x]] = tHash[t[x]] + 1
	}
	required, adquired := len(tHash), 0
	found := false
	i, j := 0, 0
	mini, minj := 0, len(s)-1
	for j < len(s) {
		sHash[s[j]] = sHash[s[j]] + 1
		if tCount, ok := tHash[s[j]]; ok {
			if tCount == sHash[s[j]] {
				adquired++
				for required == adquired && i <= j {
					found = true
					if j-i < minj-mini {
						mini, minj = i, j
					}
					sHash[s[i]] = sHash[s[i]] - 1
					if val, ok := tHash[s[i]]; ok && sHash[s[i]] < val {
						adquired--
					}
					i++
				}
			}
		}
		j++
	}
	if !found {
		return ""
	}
	return s[mini : minj+1]
}
