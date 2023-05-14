package link
/*
合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList *ListNode = &ListNode{}
	var node *ListNode = newList
	for list1 != nil && list2 != nil {
		minVal := 0
		if list1.Val <= list2.Val {
			minVal = list1.Val
			list1 = list1.Next
		} else {
			minVal = list2.Val
			list2 = list2.Next
		}
		node.Next = &ListNode{
			Val: minVal,
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}
	return newList.Next
}
