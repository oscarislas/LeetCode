//https://leetcode.com/problems/word-ladder/
//127. Word Ladder

// Given two words (beginWord and endWord), and a dictionary's word list,
// find the length of shortest transformation sequence from beginWord to endWord, such that:
//     Only one letter can be changed at a time.
//     Each transformed word must exist in the word list.
// Note:
//     Return 0 if there is no such transformation sequence.
//     All words have the same length.
//     All words contain only lowercase alphabetic characters.
//     You may assume no duplicates in the word list.
//     You may assume beginWord and endWord are non-empty and are not the same.

package main

type pair struct {
	s string
	i int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	allCombinations := make(map[string][]string, 0)
	for _, word := range wordList {
		for i := range word {
			newWord := word[:i] + "." + word[i+1:]
			allCombinations[newWord] = append(allCombinations[newWord], word)
		}
	}
	queue := []pair{
		{
			s: beginWord,
			i: 1,
		},
	}
	visited := make(map[string]struct{})
	for len(queue) > 0 {
		word := queue[0].s
		level := queue[0].i
		queue = queue[1:]

		for i := range word {
			newWord := word[:i] + "." + word[i+1:]
			combos := allCombinations[newWord]
			for _, w := range combos {
				if w == endWord {
					return level + 1
				}
				if _, ok := visited[w]; !ok {
					queue = append(queue, pair{s: w, i: level + 1})
					visited[w] = struct{}{}
				}
			}
		}
	}
	return 0
}
