/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	preNode := dummy
	for head != nil && head.Next != nil {
		first := head
		sec := head.Next
		preNode.Next = sec
		first.Next = sec.Next
		sec.Next = first

		preNode = first
		head = first.Next

	}
	return dummy.Next

}

// @lc code=end
