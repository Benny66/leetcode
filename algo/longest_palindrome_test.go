package algo

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("aca"))
}

/*
	给你一个字符串 s，找到 s 中最长的回文子串。
	如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
	示例 1：
		输入：s = "babad"
		输出："bab"
	解释："aba" 同样是符合题意的答案。
	示例 2：
		输入：s = "cbbd"
		输出："bb"
	提示：
		1 <= s.length <= 1000
		s 仅由数字和英文字母组成

解题思路：中心扩散法
从一个点向两边扩散
1、先向左边走，当左边位置等于当前位置，继续向左边走，直到不相等（left=curr）
2、然后向右边走，当右边位置等于当前位置，继续向右边走，直到不相等（right=curr）
3、判断左边位置和右边位置是否相等，相等则继续向两边走，直到不相等（left=right）
4、最后取当前左移位置到右移位置到字符串为当前点的最长回文字符串
注意：left和right分别不能小于0和超过字符串长度；
*/
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	var max int = 0
	var str string
	for cur := 0; cur < length; cur++ {
		left, right, llen, rlen := cur-1, cur+1, 0, 0
		// 1、先向左边走，当左边位置等于当前位置，继续向左边走，直到不相等（left=curr）
		for left > 0 && s[left] == s[cur] {
			left--
			//左移数+1
			llen++
		}
		// 2、然后向右边走，当右边位置等于当前位置，继续向右边走，直到不相等（right=curr）
		for right < length && s[right] == s[cur] {
			right++
			//右移数+1
			rlen++
		}
		// 3、判断左边位置和右边位置是否相等，相等则继续向两边走，直到不相等（left=right）
		for left >= 0 && right < length {
			if s[left] != s[right] {
				break
			}
			left--
			right++
			llen++
			rlen++
		}
		//长度大于最大的重新确定最长回文字符串
		if llen+rlen > max {
			max = llen + rlen
			str = s[cur-llen : cur+rlen+1]
		}
	}
	//默认第一个字符
	if max == 0 {
		str = string([]byte(s)[0])
	}
	return str
}
