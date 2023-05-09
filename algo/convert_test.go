package algo

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println(convert("abcdefg", 3))
}

type convertD struct {
	Value []string
}

func convert(s string, numRows int) string {
	for len(s)%numRows > 0 {
		s += "*"
	}
	fmt.Println(s)
	var strArr = []string{}
	m := 0
	n := numRows - 2
	l := 0
	for i := 0; i < len(s);{
		if m == 0 {
			strArr = append(strArr, string(s[i:i+numRows]))
			m++
			i += numRows 
			l = 0
		} else if m > n {
			strArr = append(strArr, string(s[i:i+numRows]))
			m = 0
			i += numRows 
		} else {
			i++
			j := n - l
			var kArr []byte
			for j > 0 {
				kArr = append(kArr, []byte("*")...)
				j--
			}
			kArr = append(kArr, s[i])
			for len(kArr) < 3 {
				kArr = append(kArr, []byte("*")...)
			}
			strArr = append(strArr, string(kArr))
			m++
			l++
		}
	}
	fmt.Println(strArr)
	return "strArr"
}
