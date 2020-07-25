package main

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	left := 0
	right := num
	middle := left + (right-left+1)/2

	for left < right {
		if middle*middle > num {
			right = middle - 1
		} else {
			left = middle
		}

		middle = left + (right-left+1)/2
	}

	return left*left == num
}

// @lc code=end
