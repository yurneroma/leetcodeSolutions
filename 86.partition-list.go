/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Val == x {
		return head
	}
	dummy := &ListNode{}
	xPoint := head
	//get x position
	for xPoint.Val != x && xPoint != nil {
		xPoint = xPoint.Next
	}

	//partition
	cur := xPoint
	for cur != nil {
		if cur.Next.Val < x {
			InsertNode(head, cur.Next)
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next

}

func InsertNode(head *ListNode, node *ListNode) {
	if node.Val < head.Val {
		node.Next = head
		head = node
		return
	}
	cur := head
	for cur != nil {
		if cur.Next.Val > node.Val {
			node.Next = cur.Next
			cur.Next = node
		} else {
			cur = cur.Next
		}
	}
}

// @lc code=end
