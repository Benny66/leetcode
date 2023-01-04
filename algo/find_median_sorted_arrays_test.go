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
