//https://leetcode.com/problems/largest-rectangle-in-histogram/
//84. Largest Rectangle in Histogram

//Given n non-negative integers representing the histogram's bar height where the width
//of each bar is 1, find the area of largest rectangle in the histogram.

package main

func largestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	positionStack := []int{0}
	maxArea := 0
	index := 1
	for index < len(heights) {
		if len(positionStack) <= 0 || heights[index] >= heights[peek(&positionStack)] {
			push(&positionStack, index)
			index++
			continue
		}
		for len(positionStack) > 0 && heights[index] < heights[peek(&positionStack)] {
			n := heights[pop(&positionStack)]
			s := index - 1 - peek(&positionStack)
			if len(positionStack) <= 0 {
				s = index
			}
			area := n * (s)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	for len(positionStack) > 0 {
		n := heights[pop(&positionStack)]
		s := index - 1 - peek(&positionStack)
		if len(positionStack) <= 0 {
			s = index
		}
		area := n * (s)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func push(stack *[]int, i int) {
	*stack = append(*stack, i)
}

func pop(stack *[]int) int {
	if len(*stack) <= 0 {
		return -1
	}
	len := len(*stack) - 1
	i := (*stack)[len]
	*stack = (*stack)[:len]
	return i
}

func peek(stack *[]int) int {
	if len(*stack) <= 0 {
		return 0
	}
	return (*stack)[len(*stack)-1]
}
