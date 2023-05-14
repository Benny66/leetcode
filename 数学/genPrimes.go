package main

import "sort"

/*
计数质数
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

*/
var prime []int

func init() {
	prime = genPrimes(5000000)
}

func genPrimes(limit int) []int {
	res := make([]int, 0)
	notPrime := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !notPrime[i] {
			res = append(res, i)
			for j := i * i; j < limit; j += i {
				notPrime[j] = true
			}
		}
	}
	return res
}

func countPrimes(n int) int {
	return sort.SearchInts(prime, n)
}
