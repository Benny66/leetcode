package algo

import (
	"fmt"
	"testing"
)

func TestMaopao(t *testing.T) {
	arr := []int{82, 3, 1, 21, 2, 2, 2, 2, 98, 34, 12, 3, 23, 4, 234, 52, 3, 42, 34, 2, 3, 41, 3, 4, 34, 5, 346, 4, 56, 45, 6, 2, 3, 42, 3, 42, 343434}
	maopao(&arr)
	fmt.Println(arr)
}

func maopao(arr *[]int) {
	for i := 0; i < len((*arr))-1; i++ {
		for j := i + 1; j < len((*arr)); j++ {
			if (*arr)[j] < (*arr)[i] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[i]
				(*arr)[i] = temp
			}
		}
	}
	return
}
