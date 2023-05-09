package array

/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

*/


func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
	//双指针，q的是填充数据的数组下标，p为覆盖数据的数组下标，
    q := 0
    p := 1
    for p < len(nums) {
		//排序的数组，当判断值不相等就表示存在不重复的数据
        if nums[p] != nums[q] {
            nums[q+1] = nums[p] 
            q++
        }
		//存在相等数值则p下标右移
        p++
    }
    return q+1
}