package main

import "fmt"

/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	fmt.Println(a, b)
	return getSum(a^b, a&b<<1)
}

// @lc code=end
