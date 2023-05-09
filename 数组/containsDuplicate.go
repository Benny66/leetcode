package array

/*
存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

示例 1：

输入：nums = [1,2,3,1]
输出：true

*/

//使用哈希map解法
func containsDuplicate(nums []int) bool {
    var tMap = make(map[int]struct{}, len(nums))
    for i:=0; i < len(nums); i++ {
        if _,ok:=tMap[nums[i]];ok{
            return true
        }
        tMap[nums[i]]=struct{}{}
    }
    return false
}