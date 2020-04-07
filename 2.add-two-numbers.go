package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	head := newNode
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		if sum >= 10 {
			carry = 1
			sum %= 10
		} else {
			carry = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		newNode.Val = sum
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		newNode.Next = &ListNode{}
		newNode = newNode.Next

	}
	if carry > 0 {
		newNode.Val = carry
	}
	return head
}

// @lc code=end
