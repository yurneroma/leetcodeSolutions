package main

import "fmt"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (28.19%)
 * Likes:    2780
 * Dislikes: 1097
 * Total Accepted:    396K
 * Total Submissions: 1.4M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list.
 *
 *
 * Note:
 *
 *
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output: 5
 *
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: 0
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 *
 *
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//edge condition
	length := len(beginWord)

	adjacencyDict := make(map[string][]string, 0)
	valList := make([]string, 0)
	//adjacency list
	for _, item := range wordList {

		for i := 0; i < length; i++ {
			key := fmt.Sprint(item[:i], "*", item[i+1:])
			valList = adjacencyDict[key]
			valList = append(valList, item)
			adjacencyDict[key] = valList
		}
	}

	//bfs
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	path := 0
	visitMap := make(map[string]bool, 0)

	for len(queue) > 0 {
		path++
		tempQueue := make([]string, 0)
		for _, word := range queue {
			for i := 0; i < length; i++ {
				key := fmt.Sprint(word[:i], "*", word[i+1:])
				for _, item := range adjacencyDict[key] {
					if item == endWord {
						return path + 1
					}
					if !visitMap[item] {
						visitMap[item] = true
						tempQueue = append(tempQueue, item)
					}
				}
			}
		}
		queue = tempQueue
	}
	return 0
}

// @lc code=end

func doubleBfs(beginWord, endWord string, wordList []string) int {

}
