//https://leetcode.com/problems/combination-sum/
//39. Combination Sum

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sums to target.
// The same repeated number may be chosen from candidates unlimited number of times.
// Note:

//     All numbers (including target) will be positive integers.
//     The solution set must not contain duplicate combinations.
// Example 1:
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]

package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	results := &[][]int{}
	combinationSumHelper(candidates, target, []int{}, results)
	return *results
}

func combinationSumHelper(candidates []int, target int, subset []int, results *[][]int) {
	for i, val := range candidates {
		newSubset := make([]int, len(subset))
		copy(newSubset, subset)
		newSubset = append(newSubset, val)
		if target-val == 0 {
			*results = append(*results, newSubset)
			break
		}
		if target-val < 0 {
			break
		}
		combinationSumHelper(candidates[i:], target-val, newSubset, results)
	}
}
