package str

/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

*/
func longestCommonPrefix(strs []string) string {
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		arr1 := strs[i]
		arr2 := str
		var p, q int = 0, 0
		for p < len(arr1) && q < len(arr2) && arr2[p] == arr1[q] {
			p++
			q++
		}
		if p == 0 {
			return ""
		}
		str = arr1[:p]
	}
	return str
}
