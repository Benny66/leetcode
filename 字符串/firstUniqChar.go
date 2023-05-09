package str
/*
字符串中的第一个唯一字符
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

示例 1：
输入: s = "leetcode"
输出: 0
*/
func firstUniqChar(s string) int {
    var hashMap = make(map[byte]struct{}, len(s))
    bytes := []byte(s)
    for i, v := range bytes {
        if _,ok:=hashMap[v];ok{
            continue
        }
        j := i+1
        for j <len(bytes) && v != bytes[j] {
            j++
        }
        if j == len(bytes){
            return i
        }
        hashMap[v] = struct{}{}
    }
    return -1
}