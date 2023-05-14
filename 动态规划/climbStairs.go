package main
/*
爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
*/
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	var array = make([]int, n+1)
	array[1] = 1
	array[2] = 2
	for i := 3; i <= n; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array[n]
}
