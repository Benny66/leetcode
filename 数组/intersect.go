package array

import "sort"

/*
两个数组的交集 II
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
*/

//排序，双指针
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var p, q int = 0, 0
	var result []int
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] == nums2[q] {
			result = append(result, nums1[p])
			p++
			q++
		} else if nums1[p] < nums2[q] {
			p++
		} else if nums1[p] > nums2[q] {
			q++
		}
	}
	return result
}
