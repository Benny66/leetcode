package array

/*
只出现一次的数字
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

示例 1 ：

输入：nums = [2,2,1]
输出：1
*/

//异或a^a^b = b a^b^a=b 所以只需要对切片内的值循环进行异或就可以得到只出现一次的值了
//位运算符
func singleNumber(nums []int) int {
    var temp int
    for i:=0;i<len(nums);i++ {
        temp ^= nums[i]
    }
    return temp
}