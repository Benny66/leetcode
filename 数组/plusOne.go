package array


/*
加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。


示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
*/


func plusOne(digits []int) []int {
    if digits[0] == 0 {
        return []int{1}
    }
    for i:=len(digits)-1;i>=0;i--{
        if digits[i]+1 != 10 {
            digits[i] = digits[i]+1
            break
        }
        digits[i] = 0
    }
    if digits[0] == 0 {
        digits = append([]int{1}, digits...)
    } 
    return digits
}