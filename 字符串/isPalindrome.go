package str

import "strings"
/*
验证回文串
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
示例 1：
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。

*/
func isPalindrome(s string) bool {
	var q, p int = 0, len(s) - 1
	for q < p {
		for q < len(s) && !('a' <= s[q] && s[q] <= 'z' || 'A' <= s[q] && s[q] <= 'Z' || '0' <= s[q] && s[q] <= '9') {
			q++
		}
		for p >= 0 && !('a' <= s[p] && s[p] <= 'z' || 'A' <= s[p] && s[p] <= 'Z' || '0' <= s[p] && s[p] <= '9') {
			p--
		}
		if q < p {
			if strings.ToLower(string(s[p])) != strings.ToLower(string(s[q])) {
				return false
			}
		}
		q++
		p--
	}
	return true
}

// func isPalindrome(s string) bool {
//     bytes := []byte(s)
//     new := []string{}
//     for _,item := range bytes {
//         if !('a' <= item && item <= 'z' || 'A' <= item && item <= 'Z' || '0' <= item && item <= '9') {
//             continue
//         }
//         new = append(new, strings.ToLower(string(item)))
//     }
//     length := len(new)
//     for i:= 0; i < length/2; i++ {
//         if new[i] != new[length-1-i] {
//             return false
//         }
//     }
//     return true

// }
