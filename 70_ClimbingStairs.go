//https://leetcode.com/problems/climbing-stairs/
//70. Climbing Stairs

// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
// Note: Given n will be a positive integer.
// Example 1:
// Input: 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps

// 1 2 3 5 8 13

// 1 1 1 1 1
// 1 1 1 2
// 1 1 2 1
// 1 2 1 1
// 2 1 1 1
// 1 2 2
// 2 1 2
// 2 2 1

package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

func climbStairsR(n int) int {
	visited := make([]int, n+1)
	visited[0], visited[1] = 1, 2
	return climbStairsHelper(n, &visited)
}

func climbStairsHelper(n int, visited *[]int) int {
	if (*visited)[n] > 0 {
		return (*visited)[n]
	}
	(*visited)[n] = climbStairsHelper(n-1, visited) + climbStairsHelper(n-2, visited)
	return (*visited)[n]
}
