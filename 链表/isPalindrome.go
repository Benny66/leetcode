package link

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
输入：head = [1,2,2,1]
输出：true
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

func init() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, math.MaxInt32)
	_out, _ := os.Create("user.out")
	out := bufio.NewWriter(_out)
next:
	for in.Scan() {
		a := in.Bytes()
		for i, n := 1, len(a); i < n/2; i++ {
			if a[i] != a[n-1-i] {
				fmt.Fprintln(out, false)
				continue next
			}
		}
		fmt.Fprintln(out, true)
	}
	out.Flush()
	os.Exit(0)
}
func isPalindrome(head *ListNode) (_bool bool) {
	return
}

// func isPalindrome(head *ListNode) bool {
//     array := []int{}
//     for head != nil {
//         array = append(array, head.Val)
//         head = head.Next
//     }
//     n := len(array)
//     for i:=0; i<n/2; i++ {
//         if array[i] != array[n-1-i] {
//             return false
//         }
//     }
//     return true
// }
// func isPalindrome(head *ListNode) bool {
//     //todo 反转
//     n := 0
//     var newHead *ListNode
//     node := head
//     for node != nil {
//         curr := &ListNode{
//             Val:node.Val,
//             Next:newHead,
//         }
//         newHead = curr
//         n++
//         node = node.Next
//     }
//     for i:=0;i<n/2;i++{
//         if newHead.Val != head.Val {
//             return false
//         }
//         newHead = newHead.Next
//         head = head.Next
//     }
//     return true
// }
