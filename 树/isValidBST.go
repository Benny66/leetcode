package tree

import "math"
/*
验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
输入：root = [2,1,3]
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

func isValidBST(root *TreeNode) bool {
	stacks := []*TreeNode{}
	min := math.MinInt64
	for len(stacks) != 0 || root != nil {
		for root != nil {
			stacks = append(stacks, root)
			root = root.Left
		}
		root = stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right

	}
	return true
}
