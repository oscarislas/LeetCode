//https://leetcode.com/problems/valid-parentheses/
//20. Valid Parentheses

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
// Example 1:
// Input: "()"
// Output: true

package main

import "errors"

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := &[]rune{}
	for _, r := range s {
		if isOpening(r) {
			pushRune(stack, r)
			continue
		}
		open, err := popRune(stack)
		if err != nil || !matches(open, r) {
			return false
		}
	}
	if len(*stack) > 0 {
		return false
	}
	return true
}

func isOpening(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

func matches(open rune, close rune) bool {
	return (open == '(' && close == ')') || (open == '{' && close == '}') || (open == '[' && close == ']')
}

func pushRune(stack *[]rune, r rune) {
	*stack = append(*stack, r)
}

func popRune(stack *[]rune) (rune, error) {
	l := len(*stack) - 1
	if l < 0 {
		return rune(0), errors.New("empty")
	}
	r := (*stack)[l]
	*stack = (*stack)[:l]
	return r, nil
}
