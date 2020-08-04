package main

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	result := ""
	if len(s) == 0 {
		return ""
	}

	repeat := 0
	bracketNumber := 0
	startIndex := 0
	for i := 0; i < len(s); i++ {
		if repeat == 0 && isNumber(s[i]) {
			for j := i; j < len(s) && isNumber(s[j]); j++ {
				repeat = repeat*10 + int(s[j]-'0')
				i = j
			}
		} else if s[i] == '[' {
			if bracketNumber == 0 {
				startIndex = i + 1
			}
			bracketNumber++
		} else if s[i] == ']' {
			bracketNumber--

			if bracketNumber == 0 {
				result = result + repeatString(decodeString(s[startIndex:i]), repeat)
				repeat = 0
				startIndex = 0
			}
		} else {
			if bracketNumber == 0 {
				result += string(s[i])
			}
		}
		// fmt.Printf("%s repeat: %d, bracketNumber: %d, startIndex:%d result: %s\n", string(s[i]), repeat, bracketNumber, startIndex, result)

	}
	return result
}

func isNumber(digit byte) bool {
	return digit >= '0' && digit <= '9'
}

func repeatString(element string, times int) string {
	result := ""

	for i := 0; i < times; i++ {
		result += element
	}

	return result
}

// @lc code=end

// func main() {
// 	fmt.Println(decodeString("3[a10[bc]]"))
// }
