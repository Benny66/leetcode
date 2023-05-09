package algo

import (
	"fmt"
	"strings"
)

func MyAtoi(s string) int {

	s = strings.TrimSpace(s)
	// Fast path for small integers that fit int type.
	s0 := s
	fmt.Println(s)
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return n
		}
		n = n*10 + int(ch)
	}
	if n > (1<<31 - 1) {
		n = 1 << 31
	}
	fmt.Println(s0[0])
	if s0[0] == '-' {
		n = -n
	}
	return n
}
