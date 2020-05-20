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

type Node struct {
	Word  string
	Level int
}

func bidirectionalBfs(beginWord, endWord string, wordList []string) int {
	adjacencyDict := make(map[string][]string, 0)
	valList := make([]string, 0)

	if !isMemBer(wordList, endWord) {
		return 0
	}
	//construct adjacency list
	for _, word := range wordList {
		for i := 0; i < len(beginWord); i++ {
			key := fmt.Sprint(word[:i], "*", word[i+1:])
			valList = adjacencyDict[key]
			valList = append(valList, word)
			adjacencyDict[key] = valList
		}
	}

	//double bfs
	queueBegin := make([]Node, 0)
	queueEnd := make([]Node, 0)
	queueBegin = append(queueBegin, Node{Word: beginWord, Level: 1})
	queueEnd = append(queueEnd, Node{Word: endWord, Level: 1})
	visitedBegin := make(map[string]int, 0)
	visitedEnd := make(map[string]int, 0)
	visitedBegin[beginWord] = 1
	visitedEnd[endWord] = 1
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		path, tempQueueBegin := visitNode(queueBegin, visitedBegin, visitedEnd, adjacencyDict)
		if path > -1 {
			fmt.Println("path ", path)
			return path
		}

		path, tempQueueEnd := visitNode(queueEnd, visitedEnd, visitedBegin, adjacencyDict)
		if path > -1 {
			fmt.Println("path ", path)
			return path
		}
		queueBegin = tempQueueBegin
		queueEnd = tempQueueEnd
	}
	return 0
}

func visitNode(queue []Node, visited, otherVisited map[string]int, adjacencyDict map[string][]string) (int, []Node) {
	tempQueue := make([]Node, 0)
	level := -1
	for _, item := range queue {
		level = item.Level
		for i := 0; i < len(item.Word); i++ {
			key := fmt.Sprint(item.Word[:i], "*", item.Word[i+1:])
			valList := adjacencyDict[key]
			for _, word := range valList {
				if otherVisited[word] > 0 {
					fmt.Println("the level is ", level)
					fmt.Println("valList", valList)
					fmt.Println("other", otherVisited)
					return level + otherVisited[word], nil
				}
				if visited[word] == 0 {
					visited[word] = level + 1
					tempQueue = append(tempQueue, Node{Level: level + 1, Word: word})
				}
			}
		}
	}
	return -1, tempQueue
}

func isMemBer(s []string, item string) bool {
	for _, elem := range s {
		if elem == item {
			return true
		}
	}
	return false
}

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	path := bidirectionalBfs(beginWord, endWord, wordList)
	return path
}

// @lc code=end

func bfs(beginWord, endWord string, wordList []string) int {
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

// func main() {
// 	beginWord := "a"
// 	endWord := "c"
// 	wordList := []string{"a", "b", "c"}
// 	fmt.Println(DoubleBfs(beginWord, endWord, wordList))
// }
