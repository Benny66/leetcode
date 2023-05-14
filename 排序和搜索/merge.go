package levelOrder

/*
合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int = m - 1, n - 1
	var end = m + n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[end] = nums1[i]
			i--
		} else {
			nums1[end] = nums2[j]
			j--
		}
		end--
	}
}

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     var array = make([]int, m+n)
//     var index, q, p int = 0, 0, 0
//     for q < m && p < n {
//         if nums1[q] <= nums2[p] {
//             array[index] = nums1[q]
//             q++
//         }else{
//             array[index] = nums2[p]
//             p++
//         }
//         index++
//     }
//     for q < m {
//         array[index] = nums1[q]
//         index++
//         q++
//     }
//     for p < n {
//         array[index] = nums2[p]
//         index++
//         p++
//     }
//     fmt.Println(array)
//     for i:=0; i<len(array);i++ {
//         nums1[i] = array[i]
//     }
// }
