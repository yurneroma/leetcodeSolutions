/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//要用当前节点，pre 比较令人困惑。 cur = cur.Next equal  pre = cur
	cur := head
	for cur.Next != nil && cur != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	// pre := head
	// cur := head.Next
	// for cur != nil {
	// 	if pre.Val == cur.Val {
	// 		pre.Next = cur.Next
	// 	} else {
	// 		pre = cur
	// 	}
	// 	cur = cur.Next
	// }
	return head
}

// @lc code=end

