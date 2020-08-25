import (
	"strconv"
)

/*
* @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
*/

// @lc code=start
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	var i = 0
	for j := i + 1; j < len(num); j++ {
		for k := j + 1; k < len(num); k++ {
			if dfs(&num, i, j, k) {
				return true
			}
		}
	}

	return false
}

//       i j k l
// num = 123456
// number1 = 12
// number2 = 34
func dfs(num *string, i, j, k int) bool {

	if k == len(*num) {
		return true
	}

	number1 := (*num)[i:j]
	number2 := (*num)[j:k]

	if len(number1) != 1 && number1[0] == '0' {
		return false
	}
	if len(number2) != 1 && number2[0] == '0' {
		return false
	}

	for l := k + 1; l <= len(*num); l++ {
		number3 := (*num)[k:l]
		if additive(number1, number2) == number3 {
			return dfs(num, j, k, l)
		}
	}
	return false
}

func additive(num1, num2 string) string {
	var longer, shorter string
	if len(num1) > len(num2) {
		longer, shorter = num1, num2
	} else {
		longer, shorter = num2, num1
	}

	var result = ""
	var x = 0
	var y = 0
	var shorterDigit = 0
	var longerDigit = 0
	for i := 1; i <= len(longer); i++ {

		if i <= len(shorter) {
			shorterDigit, _ = strconv.Atoi(string(shorter[len(shorter)-i]))
		} else {
			shorterDigit = 0
		}

		longerDigit, _ = strconv.Atoi(string(longer[len(longer)-i]))
		current := longerDigit + shorterDigit + x
		x = current / 10
		y = current % 10
		result = strconv.Itoa(y) + result
	}

	if x != 0 {
		result = strconv.Itoa(x) + result
	}
	return result
}

// @lc code=end
