package str

import "math"
/*
整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
示例 1：
输入：x = 123
输出：321
*/
func reverse(x int) int {
	var res int
	for x != 0 {
		t := x % 10
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		res = res*10 + t
		x = x / 10
	}
	return res
}

// func reverse(x int) int {
//     var index int
//     var arr []int
//     for x/10 != 0 {
//         arr = append(arr, x%10)
//         x = x/10
//         index++
//     }
//     arr = append(arr, x)
//     var res int
//     for key,item:=range arr{
//         if key == 0 && item == 0 {
//             index--
//             continue
//         }
//         res += item * int(math.Pow(10,float64(index)))
//         index--
//     }
//     if res < -(1 << 31) || res > (1 << 31 -1) {
//         return 0
//     }
//     return res
// }
