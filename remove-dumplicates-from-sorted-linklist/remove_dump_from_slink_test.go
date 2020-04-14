package RemoveSlink

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 3}
	l5 := &ListNode{Val: 3}
	l6 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	head := DeleteDuplicates(l1)

	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}
