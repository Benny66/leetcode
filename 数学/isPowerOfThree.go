package main

import "math"

/*
3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
示例 1：
输入：n = 27
输出：true

*/
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	//具体来说，当计算math.Log(float64(27)) / math.Log(3)时，结果应该是3.0，但由于浮点数运算的精度限制，实际计算结果可能略微偏离这个值。例如，如果计算结果为3.0000000000000004，则小数部分不等于0，isPowerOfThree()函数将返回false。
	log := math.Log(float64(n)) / math.Log(3)
	x := int(math.Round(log))

	return int(math.Pow(3, float64(x))) == n
}
