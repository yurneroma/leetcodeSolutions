/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var s [][]int
	slice := make([]*TreeNode, 0)
	slice = append(slice, root)
	left2Right := false
	for len(slice) != 0 {
		tmpSlice := make([]*TreeNode, 0)
		valList := make([]int, 0)
		for _, node := range slice {
			valList = append(valList, node.Val)
			if node.Left != nil {
				tmpSlice = append(tmpSlice, node.Left)
			}
			if node.Right != nil {
				tmpSlice = append(tmpSlice, node.Right)
			}
		}
		slice = tmpSlice
		if left2Right {
			reverse(valList)
		}
		left2Right = !left2Right
		s = append(s, valList)
	}
	return s
}

func reverse(val []int) {
	if len(val) == 0 {
		return
	}

	for i := 0; i < len(val)/2; i++ {
		val[i], val[len(val)-1-i] = val[len(val)-1-i], val[i]
	}
}

// @lc code=end

