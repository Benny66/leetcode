package str

/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
示例 1:
输入: s = "anagram", t = "nagaram"
输出: true
*/
func isAnagram(s string, t string) bool {
 
    if len(s)!=len(t){
        return false
    }
    ss:=make([]int,26)
    tt:=make([]int,26)

    for i,_:=range s{
        ss[s[i]-'a']++
    }
    for i,_:=range t{
        tt[t[i]-'a']++
    }
   
    for i:=0;i<26;i++{
        if ss[i]!=tt[i]{
            return false
        }
    }
    return true

}
// func isAnagram(s string, t string) bool {
//     arrS :=[]byte(s)
//     arrT := []byte(t)
//     sort.Slice(arrS, func(i, j int) bool {
// 		return arrS[i] < arrS[j]
// 	})
//     sort.Slice(arrT, func(i, j int) bool {
// 		return arrT[i] < arrT[j]
// 	})
//     return string(arrT) == string(arrS)
// }
// func isAnagram(s string, t string) bool {
//     arrS := []byte(s)
//     arrT := []byte(t)
//     var hashMap = make(map[byte]int, len(s)) 
//     for _,item:=range arrS {
//         if _,ok:=hashMap[item];ok{
//             hashMap[item]++
//         }else{
//             hashMap[item] = 1
//         }
//     }
//     for _,item:=range arrT {
//         if _,ok:=hashMap[item];!ok{
//             return false
//         }
//         hashMap[item] -= 1
//         if hashMap[item] == 0 {
//             delete(hashMap, item)
//         }
//     }
//     if len(hashMap) > 0 {
//         return false
//     }
//     return true
// }