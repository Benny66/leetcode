package algo

import (
	"fmt"
	"testing"
)

func TestContinuousNumMax(t *testing.T) {
	result := ContinuousNumMax([]int{1, 4, 1, 2, 4, 5, 3, 2, 2, 5, 100, 2}, 3)
	fmt.Println(result)

	result = ContinuousNumMax2([]int{1, 4, 1, 2, 4, 5, 3, 2, 2, 5, 100, 2}, 3)
	fmt.Println(result)
}

// 穷举
func ContinuousNumMax(array []int, n int) int {
	if len(array) <= n {
		return sumArray(array)
	}
	var max int
	var right int = n
	for left := 0; left <= len(array)-n; left++ {
		fmt.Println(array[left:right])
		max = maxInt(max, sumArray(array[left:right]))
		right++
	}
	return max
}

// 滑动窗口
func ContinuousNumMax2(array []int, n int) int {
	var max int
	var sum int
	var num int = 0
	for i := 0; i < len(array); i++ {
		if num == n {
			sum = sum - array[i-n] + array[i]
			max = maxInt(max, sum)
			continue
		}
		sum += array[i]
		num++
		max = sum
	}
	return max
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func sumArray(array []int) int {
	var result int
	for i := 0; i < len(array); i++ {
		result += array[i]
	}
	return result
}
