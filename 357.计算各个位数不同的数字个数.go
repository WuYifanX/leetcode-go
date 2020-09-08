// package main

/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

//  f(0)=1
// f(1)=10
// f(2)=9*9+f(1)
// f(3)=9*9*8+f(2)
// f(4)=9*9*8*7+f(3)

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}

	uniqueDigits := 1

	for i := 0; i < n; i++ {
		if i == 0 {
			uniqueDigits *= 9
		} else {
			uniqueDigits *= (10 - i)
		}
	}

	return uniqueDigits + countNumbersWithUniqueDigits(n-1)

}

// @lc code=end
