package algo

import "fmt"

func GetNStr(str string, num int) string {
	if num == 1 {
		return str
	}
	bytes := []byte(str)
	fmt.Println(bytes)
	var sts = make([][]byte, 0)
	var temp = make([]byte, 0)
	var mNum = num - 2
	var mNumM = 0
	//一个周期，是多少个字符
	var cycle = num + mNum
	for i := 0; i < len(bytes); i++ {
		//周期内小于num则为同一行
		if i%cycle < num {
			temp = append(temp, bytes[i])
			if len(temp) == num {
				sts = append(sts, temp)
				temp = []byte{}
			}
		} else {
			//中间几行
			temp = append(temp, make([]byte, mNum-mNumM)...)
			temp = append(temp, bytes[i])
			temp = append(temp, make([]byte, num-len(temp))...)
			sts = append(sts, temp)
			temp = []byte{}
			mNumM++
			//判断中间是否达上限
			if mNumM >= mNum {
				mNumM = 0
			}
		}
	}
	if len(temp) != 0 {
		sts = append(sts, append(temp, make([]byte, num-len(temp))...))
	}
	fmt.Println(sts, len(sts))
	var result string
	for i := 0; i < num; i++ {
		for j := 0; j < len(sts); j++ {
			if sts[j][i] != 0 {
				result += string(sts[j][i])
			}
		}
	}
	fmt.Println(result)
	return result
}

/* ABCDEFGHIJKLMM
3 14/3 4.2
4 14/4 3.2
5 14/5 2.4
*/

func GetNStr2(str string, num int) string {
	if num < 2 {
		return str
	}
	var index int
	var arrays = make(map[int]string, 0)
	var isEnd bool = true
	for i := 0; i < len(str); i++ {
		fmt.Println(index)
		arrays[index] += string(str[i])
		if index < num-1 && isEnd {
			index++
		} else {
			isEnd = false
			index--
			if index == 0 {
				isEnd = true
			}
		}
	}
	fmt.Println(arrays)
	var result string
	for i := 0; i < num; i++ {
		result += arrays[i]
	}
	return result
}
