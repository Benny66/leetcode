package tree
/*
对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
输入：root = [1,2,2,3,4,4,3]
输出：true
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isHandleSymmetric(root.Left, root.Right)
}
func isHandleSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isHandleSymmetric(left.Left, right.Right) && isHandleSymmetric(left.Right, right.Left)
}
