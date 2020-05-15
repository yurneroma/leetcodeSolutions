/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (52.91%)
 * Likes:    2632
 * Dislikes: 65
 * Total Accepted:    569.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var s [][]int
	slice := make([]*TreeNode, 0)
	slice = append(slice, root)
	for len(slice) != 0 {
		tempSlice := make([]*TreeNode, 0)
		valList := make([]int, 0)
		for _, node := range slice {
			valList = append(valList, node.Val)
			if node.Left != nil {
				tempSlice = append(tempSlice, node.Left)
			}
			if node.Right != nil {
				tempSlice = append(tempSlice, node.Right)
			}
		}
		s = append(s, valList)
		slice = tempSlice

	}
	return s
}

// @lc code=end

