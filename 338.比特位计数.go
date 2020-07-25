package main

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	maps := make(map[int]int)
	maps[0] = 0
	maps[1] = 1
	for i := 1; i <= 31; i++ {
		maps[2<<(i-1)] = i
	}

	result := make([]int, num+1)
	digit := 0
	result[0] = 0
	for i := 1; i <= num; i++ {
		currentDigit, isExist := maps[i]

		if isExist {
			digit = currentDigit
			result[i] = 1
		} else {
			result[i] = result[2<<(digit-1)] + result[i-2<<(digit-1)]
		}
	}
	return result
}

// @lc code=end
