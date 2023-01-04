package algo

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	//abcaabcdecd
	fmt.Println(lengthOfLongestSubstring("abcaabcdecd"))
}

func lengthOfLongestSubstring(s string) int {
	var maxLen int = 0
	if s == "" {
		return maxLen
	}
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		slice := bytes[i:]
		var sMap = make(map[byte]bool, len(slice))
		for _, item := range slice {
			if _, ok := sMap[item]; ok {
				break
			}
			sMap[item] = true
		}
		if maxLen < len(sMap) {
			maxLen = len(sMap)
		}
	}
	return maxLen
}
