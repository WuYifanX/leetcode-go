package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.85%)
 * Likes:    495
 * Dislikes: 0
 * Total Accepted:    200.8K
 * Total Submissions: 516.8K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	if x <= 3 {
		return 1
	}

	left, right := 1, (x+1)/2

	var middle int

	for left < right {
		middle = left + (right-left+1)/2

		if middle*middle > x {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return left
}

// @lc code=end
