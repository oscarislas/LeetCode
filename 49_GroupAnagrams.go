//https://leetcode.com/problems/group-anagrams/
//49. Group Anagrams

// Given an array of strings, group anagrams together.
// Example:
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:
//     All inputs will be in lowercase.
//     The order of your output does not matter.

package main

import (
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	m := make(map[string][]string)
	for _, word := range strs {
		count := make([]int, 26)
		for _, r := range word {
			index := int(r) - int('a')
			count[index]++
		}
		strIndex := ""
		for _, c := range count {
			strIndex = strIndex + "#" + strconv.Itoa(c)
		}
		m[strIndex] = append(m[strIndex], word)
	}
	for _, val := range m {
		result = append(result, val)
	}
	return result
}
