package main

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (50.63%)
 * Likes:    1161
 * Dislikes: 203
 * Total Accepted:    297K
 * Total Submissions: 582.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var val [][]int
	walkByDFS(root, 0, &res)
	for i := len(res) - 1; i >= 0; i-- {
		slice := res[i]
		val = append(val, slice)
	}
	return val
}

//DFS
func walkByDFS(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if level == len(*res) {
		list := make([]int, 0)
		*res = append(*res, list)
	}
	elemList := (*res)[level]
	elemList = append(elemList, root.Val)
	(*res)[level] = elemList
	walkByDFS(root.Left, level+1, res)
	walkByDFS(root.Right, level+1, res)

}

//队列存层节点， 遍历层节点，val 按层放入二维数组，  最后二维数组倒序输出
func walkByBFS(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	node := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for len(queue) != 0 {
		level := make([]int, 0)
		count := len(queue)
		for i := 0; i < count; i++ {
			subNode := queue[i]
			level = append(level, subNode.Val)
			if subNode.Left != nil {
				queue = append(queue, subNode.Left)
			}
			if subNode.Right != nil {
				queue = append(queue, subNode.Right)
			}

		}
		res = append(res, level)
		queue = queue[count:]

	}
	val := make([][]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		val = append(val, res[i])
	}
	return val

}

// @lc code=end
