package array
/*
旋转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

*/
//使用临时数组
func rotate(nums []int, k int) {
    temp := make([]int, len(nums))
    copy(temp, nums)
    length := len(nums)
    for i:= 0;i < length; i++ {
        nums[(i+k)%length] = temp[i]
    }
}