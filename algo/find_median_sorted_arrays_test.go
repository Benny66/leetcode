package algo

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{2, 3}, []int{1, 4}))
}

/*
	题目：寻找两个正序数组的中位数
	给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
	算法的时间复杂度应该为 O(log (m+n)) 。
	示例 1：
		输入：nums1 = [1,3], nums2 = [2]
		输出：2.00000
		解释：合并数组 = [1,2,3] ，中位数 2
	示例 2：
		输入：nums1 = [1,2], nums2 = [3,4]
		输出：2.50000
		解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
	提示：
		nums1.length == m
		nums2.length == n
		0 <= m <= 1000
		0 <= n <= 1000
		1 <= m + n <= 2000
		-106 <= nums1[i], nums2[i] <= 106

解题思路：合并两个数组，返回排序，使用了golang内置的sort排序
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	hNum := append(nums1, nums2...)
	sort.Ints(hNum)
	fmt.Println(hNum)
	hLen := len(hNum)
	var result float64
	if hLen%2 == 0 {
		result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", (float64(hNum[hLen/2])+float64(hNum[hLen/2-1]))/2), 64)
	} else {
		result, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(hNum[hLen/2])), 64)
	}
	return result
}
