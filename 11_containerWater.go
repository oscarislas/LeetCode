//https://leetcode.com/problems/container-with-most-water/
// 11. Container With Most Water

// Given n non-negative integers a1, a2, ..., an , where each represents a point
// at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
// of line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container,
// such that the container contains the most water.
// Note: You may not slant the container and n is at least 2.

package main

func maxArea(height []int) int {
	x := 0
	y := len(height) - 1
	maxArea := 0
	if y > 0 {
		for {
			area := 0
			if height[x] < height[y] {
				area = height[x] * (y - x)
				x++
			} else {
				area = height[y] * (y - x)
				y--
			}
			if maxArea < area {
				maxArea = area
			}
			if x >= y {
				break
			}
		}
	}
	return maxArea
}
