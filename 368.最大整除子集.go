// package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	sort.Ints(nums)

	dp := make([]int, len(nums))
	maxLength := 1
	for i := 0; i < len(nums); i++ {
		currentMax := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				currentMax = max(dp[j], currentMax)
			}
		}

		dp[i] = currentMax + 1
		maxLength = max(maxLength, dp[i])
	}
	// fmt.Println(dp, maxLength)
	result := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		// fmt.Println(dp[i], maxLength, nums[i])
		if dp[i] == maxLength && maxLength != 0 {
			if len(result) == 0 {
				result = append(result, nums[i])
				maxLength--
			} else if result[len(result)-1]%nums[i] == 0 {
				result = append(result, nums[i])
				maxLength--
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

// func main() {
// 	fmt.Println(largestDivisibleSubset([]int{2, 3, 4, 8}))
// }
