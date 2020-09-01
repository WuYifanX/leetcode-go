package main

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (40.92%)
 * Likes:    791
 * Dislikes: 0
 * Total Accepted:    127.9K
 * Total Submissions: 310.8K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 *
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 *
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = 100000
		for coinIndex := 0; coinIndex < len(coins); coinIndex++ {
			coin := coins[coinIndex]
			if i-coin >= 0 {
				dp[i] = minOf(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] == 100000 {
		return -1
	}
	return dp[amount]
}

func minOf(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
