//https://leetcode.com/problems/valid-palindrome/
//125. Valid Palindrome

// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// Note: For the purpose of this problem, we define empty string as valid palindrome.
// Example 1:
// Input: "A man, a plan, a canal: Panama"
// Output: true
package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l := len(s)
	mid := int(l / 2)
	for i, j := 0, l-1; i < mid; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func recursiveIsPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l := len(s)
	if l <= 0 {
		return true
	}
	if l == 1 {
		return true
	}
	if s[0] != s[l-1] {
		return false
	}
	return isPalindrome(s[1 : l-1])
}
