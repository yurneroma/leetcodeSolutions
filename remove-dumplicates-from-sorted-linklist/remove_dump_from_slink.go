package RemoveSlink

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
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
	return head
}
