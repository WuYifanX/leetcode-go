// package main

/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 *
 * https://leetcode-cn.com/problems/integer-replacement/description/
 *
 * algorithms
 * Medium (35.37%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 21.6K
 * Testcase Example:  '8'
 *
 * 给定一个正整数 n，你可以做如下操作：
 *
 * 1. 如果 n 是偶数，则用 n / 2替换 n。
 * 2. 如果 n 是奇数，则可以用 n + 1或n - 1替换 n。
 * n 变为 1 所需的最小替换次数是多少？
 *
 * 示例 1:
 *
 *
 * 输入:
 * 8
 *
 * 输出:
 * 3
 *
 * 解释:
 * 8 -> 4 -> 2 -> 1
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * 7
 *
 * 输出:
 * 4
 *
 * 解释:
 * 7 -> 8 -> 4 -> 2 -> 1
 * 或
 * 7 -> 6 -> 3 -> 2 -> 1
 *
 *
 */

// @lc code=start
func integerReplacement(n int) int {

	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return minOf(integerReplacement(n+1), integerReplacement(n-1)) + 1
	}

}

func minOf(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
