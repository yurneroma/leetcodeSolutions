/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	tail := head
	leng := 1
	for tail.Next != nil {
		tail = tail.Next
		leng++
	}

	tail.Next = head
	tail = head
	for i := 1; i < leng-k%leng; i++ {
		tail = tail.Next
	}
	head, tail.Next = tail.Next, nil

	return head

}

// @lc code=end

