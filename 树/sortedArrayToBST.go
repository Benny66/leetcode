package tree
/*
将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

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

/*
判断数组是否为空，如果为空则返回 null。
找到数组的中间位置 mid。
创建一个新的二叉树节点，值为 nums[mid]。
以 mid 为界，将数组分为左右两个部分 leftNums 和 rightNums。
将左子树构建为 sortedArrayToBST(leftNums)。
将右子树构建为 sortedArrayToBST(rightNums)。
将当前节点的左子树和右子树分别设置为步骤 5 和步骤 6 中构建的子树。
返回当前节点。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
