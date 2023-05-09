package algo

import (
	"fmt"
	"testing"
)

func TestContinuousSumResult(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61}
	result := ContinuousSumResult(array, 200)
	fmt.Println(result)
	result = ContinuousSumResult2(array, 200)
	fmt.Println(result)
}

// 穷举
func ContinuousSumResult(array []int, data int) [][]int {
	var left int = 0
	var length = len(array)
	var resultArr [][]int
	for left < length {
		var cResultArr []int
		var cSum int = 0
		for i := left; i < len(array); i++ {
			cSum += array[i]
			if cSum > data {
				break
			}
			cResultArr = append(cResultArr, array[i])
			if cSum == data {
				resultArr = append(resultArr, cResultArr)
				break
			}
		}
		left++
	}

	return resultArr
}

// 滑动
// 1, 2, 3, 1, 2, 3
func ContinuousSumResult2(array []int, data int) [][]int {
	var left, right int     //左、右窗口指针下标
	var length = len(array) //数组长度
	var resultArr [][]int   //保存的结果
	var list []int          //窗口列表
	var sum int = 0         //窗口列表的和
	for left < length && right < length {
		for right < length {
			//当前窗口的数组和大于，则左窗口指针右移
			if sum+array[right] > data {
				//窗口列表无数据无法扣减
				if len(list) != 0 {
					sum -= list[0]
					list = list[1:]
				}
				left++
				break
			}
			//小于或者等于则数值添加，推到list
			sum += array[right]
			list = append(list, array[right])
			//当前窗口的数组和等于，则左、右窗口指针右移
			if sum == data {
				//保存结果
				resultArr = append(resultArr, list)
				//窗口列表无数据无法扣减
				if len(list) != 0 {
					sum -= list[0]
					list = list[1:]
				}
				left++
				right++
				break
			}
			//当前窗口的数组和小于，则右窗口指针右移
			right++
		}
		//当左窗口指针大于右窗口，交换指针下标
		if left > right {
			temp := right
			right = left
			left = temp
		}
	}

	return resultArr
}
