package algo

import (
	"fmt"
	"strconv"
)

func GetIntF(data int) int {
	var s0 bool
	if data <= 0 {
		data = -data
		s0 = true
	}
	str := strconv.Itoa(data)
	bytes := []byte(str)
	j := len(bytes) - 1
	for i := 0; i < len(bytes)/2; i++ {
		temp := bytes[i]
		bytes[i] = bytes[j]
		bytes[j] = temp
		j--
	}
	result, err := strconv.Atoi(string(bytes))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if result-1 > (1 << 31) {
		return 0
	}
	if s0 {
		result = -result
	}

	return result
}
