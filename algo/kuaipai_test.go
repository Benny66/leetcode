package algo

import (
	"fmt"
	"testing"
)

func TestKuaiPai(t *testing.T) {
	arr := []int{82,3,1,21,2,2,2,2,98,34,12,3,23,4,234,52,3,42,34,2,3,41,3,4,34,5,346,4,56,45,6,2,3,42,3,42,343434}
	kuaipai(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func kuaipai(arr *[]int, start int, end int) {
	if start > end {
		return
	}
	value := (*arr)[start]
	lKey := start
	rKey := end
	for start != end {
		for start < end && (*arr)[end] >= value {
			end--
		}
		for start < end && (*arr)[start] <= value {
			start++
		}
		if start < end {
			temp := (*arr)[end]
			(*arr)[end] = (*arr)[start]
			(*arr)[start] = temp
		}
	}
	(*arr)[lKey] = (*arr)[start]
	(*arr)[start] = value
	kuaipai(arr, lKey, start-1)
	kuaipai(arr, start+1, rKey)
}
