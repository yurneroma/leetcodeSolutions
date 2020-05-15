/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (36.83%)
 * Likes:    1216
 * Dislikes: 646
 * Total Accepted:    393K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 * Given binary tree [3,9,20,null,null,15,7],
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * return its minimum depth = 2.
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
func minDepth(root *TreeNode) int {

	//bfs
	//return bfs(root)

	//recursive
	//return recursive(root)

	//dfs
	return dfsSolution(root)
}

//bfs solution
func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var mdepth int
	queue := make([]*TreeNode, 0)
	mdepth = 1
	queue = append(queue, root)
re:
	for len(queue) != 0 {
		tempQueue := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				break re
			}
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}
		mdepth++
		queue = tempQueue
	}
	return mdepth
}

//recursive solution
func recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	m1 := recursive(root.Left)
	m2 := recursive(root.Right)
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	}
	return intMin(m1, m2) + 1
}
func intMin(m1, m2 int) int {
	if m1 > m2 {
		m1 = m2
	}
	return m1
}

var min int

//dfs solution
func dfsSolution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min = 1<<32 - 1
	dfs(root, 1)
	return min
}

func dfs(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		if depth < min {
			min = depth
		}
	}

	if root.Left != nil {
		dfs(root.Left, depth+1)
	}
	if root.Right != nil {
		dfs(root.Right, depth+1)
	}
}

// @lc code=end

