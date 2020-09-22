package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	intiallen, arrlen := len(strs[0]), len(strs)
	str := []byte{}
	for i := 0; i < intiallen; i++ {
		for j := 1; j < arrlen; j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				return string(str)
			}
		}
		str = append(str, strs[0][i])

	}
	return string(str)
}

func main() {
	fmt.Print(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Print(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
