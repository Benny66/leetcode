package algo

import "testing"

func TestContinuousLengthMin(t *testing.T) {

}

// {1,7,4,3,2,3,1,1,1}, //5
func ContinuousLengthMin(array []int, data int) []int {

	var left, right int
	var length int = len(array)
	var sum int = 0
	var result []int
	var list []int
	var minListNum int = length
	for left < length && right < length {
		for right < length && len(list) < minListNum {
			if sum+array[right] > data {
				if len(list) != 0 {
					sum -= list[0]
					list = list[1:]
				}
				left++
			}
			sum += array[right]
			list = append(list, array[right])
			if sum+array[right] == data {
				minListNum = len(list)
				if len(list) != 0 {
					sum -= list[0]
					list = list[1:]
				}
				left++
				right++
			}
			right++
		}
		if left > right {
			temp := right
			right = left
			left = temp
		}
	}

	return result
}
