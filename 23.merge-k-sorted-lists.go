/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	leng := len(lists)
	if leng == 0 {
		return nil
	}
	for leng > 1 {

		for i := 0; i < leng/2; i++ {
			lists[i] = merge2Lists(lists[i], lists[leng-i-1])
		}
		leng = (leng + 1) / 2
	}
	return lists[0]

}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	//l1,l2 任一个遍历到tail，都跳出循环。
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}

	//cur 节点，指向长度更长的那个链表。
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

// @lc code=end

