package link

/*
反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
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

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	//  var currHead *ListNode
	for head != nil {
		var currHead *ListNode = &ListNode{
			Val:  head.Val,
			Next: newHead,
		}
		newHead = currHead
		head = head.Next
	}
	return newHead
}

// func reverseList(head *ListNode) *ListNode {
//     array := []int{}
//     var newHead *ListNode = &ListNode{}
//     for head != nil {
//         array = append(array, head.Val)
//         head = head.Next
//     }
//     if len(array) == 0 {
//         return nil
//     }
//     node := newHead
//     for i:=len(array)-1;i>=0;i--{
//         node.Val = array[i]
//         if i != 0 {
//             node.Next = &ListNode{}
//             node = node.Next
//         }
//     }
//     return newHead
// }
