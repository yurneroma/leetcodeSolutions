/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	tail := head
	dummy := &ListNode{}
	dummy.Next = head
	//get tail of the LinkList
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	//必须要有哨兵节点， 以免删除head 的情况
	pre := dummy
	for i := 0; i < (length - n); i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next

}

// @lc code=end

