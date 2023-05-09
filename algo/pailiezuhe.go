package algo

func GetP(str string) []string {
	var result []string
	for i := 0; i < len(str); i++ {
		result = zuHe(result, initData[string(str[i])])
	}
	return result
}

var initData = map[string][]string{
	"2": {
		"a", "b", "c",
	},
	"3": {
		"d", "e", "f",
	},
	"4": {
		"g", "h", "i",
	},
	"5": {
		"j", "k", "l",
	},
	"6": {
		"m", "n", "o",
	},
	"7": {
		"p", "q", "r", "s",
	},
	"8": {
		"t", "u", "v",
	},
	"9": {
		"w", "x", "y", "z",
	},
}

func zuHe(arr1, arr2 []string) (result []string) {
	if len(arr1) == 0 {
		result = arr2
		return
	}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			result = append(result, arr1[i]+arr2[j])
		}
	}
	return
}
