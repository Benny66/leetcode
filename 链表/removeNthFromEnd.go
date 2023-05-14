package link
/*
删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。


*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	m := 0
	node := head
	for node != nil {
		node = node.Next
		m++
	}
	j := m - n
	if j == 0 {
		return head.Next
	}
	node = head
	for i := 0; i < j-1; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return head
}
