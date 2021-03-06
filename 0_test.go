package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotatedSortedArray(t *testing.T) {
	assert.Equal(t, -1, search(nil, 0))
	assert.Equal(t, -1, search([]int{}, 0))
	assert.Equal(t, -1, search([]int{1}, 0))
	assert.Equal(t, 0, search([]int{1}, 1))
	assert.Equal(t, -1, search([]int{1}, 2))
	assert.Equal(t, 1, search([]int{3, 1}, 1))
	assert.Equal(t, 0, search([]int{3, 1}, 3))
	assert.Equal(t, 1, search([]int{1, 3}, 3))
	assert.Equal(t, 2, search([]int{5, 1, 3}, 3))
	assert.Equal(t, 0, search([]int{1, 3}, 1))
	assert.Equal(t, -1, search([]int{1, 3}, 4))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}

func TestMergeTwoList(t *testing.T) {
	assert.Equal(t, []int{}, mergeTwoListsHelper(nil, nil))
	assert.Equal(t, []int{1, 2, 4}, mergeTwoListsHelper([]int{1, 2, 4}, nil))
	assert.Equal(t, []int{1, 3, 4}, mergeTwoListsHelper(nil, []int{1, 3, 4}))
	assert.Equal(t, []int{1, 2}, mergeTwoListsHelper([]int{2}, []int{1}))
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, mergeTwoListsHelper([]int{1, 2, 4}, []int{1, 3, 4}))
}

func TestValidParentheses(t *testing.T) {
	assert.Equal(t, true, isValid(""))
	assert.Equal(t, false, isValid("((("))
	assert.Equal(t, false, isValid("(("))
	assert.Equal(t, true, isValid("(){}[]"))
	assert.Equal(t, false, isValid("([)]"))
	assert.Equal(t, true, isValid("([{}])"))
}

func TestThreeSum(t *testing.T) {
	assert.Equal(t, [][]int{}, threeSum([]int{}))
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0, 0, 0}))
}
func TestContainerWater(t *testing.T) {
	assert.Equal(t, 0, maxArea([]int{}))
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 6, maxArea([]int{1, 2, 3, 4, 5}))
}

func TestLongesSubstring(t *testing.T) {
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 2, lengthOfLongestSubstring("aab"))
	assert.Equal(t, 4, lengthOfLongestSubstring("abcda"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdfdd"))
}

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "aaaa", longestPalindrome("aaaa"))
	assert.Equal(t, "a", longestPalindrome("abc"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 0, largestRectangleArea([]int{}))
	assert.Equal(t, 20, largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0}))
	assert.Equal(t, 6, largestRectangleArea([]int{4, 2, 0, 3, 2, 5}))
	assert.Equal(t, 6, largestRectangleArea([]int{1, 2, 3, 2, 1}))
	assert.Equal(t, 7, largestRectangleArea([]int{1, 3, 4, 1, 1, 1, 1}))
	assert.Equal(t, 16, largestRectangleArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 2, largestRectangleArea([]int{1, 1}))
	assert.Equal(t, 1, largestRectangleArea([]int{1}))
	assert.Equal(t, 9, largestRectangleArea([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 9, largestRectangleArea([]int{1, 2, 3, 4, 5}))
}

func TestValidPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("aaaazzaaaa"))
	assert.Equal(t, true, isPalindrome("aba"))
	assert.Equal(t, true, isPalindrome("a"))
	assert.Equal(t, false, isPalindrome("zs"))
	assert.Equal(t, false, isPalindrome("zszf"))
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
}
