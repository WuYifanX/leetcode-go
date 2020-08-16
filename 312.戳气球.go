/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 *
 * https://leetcode-cn.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (61.31%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 40.6K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 *
 * 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的
 * left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
 *
 * 求所能获得硬币的最大数量。
 *
 * 说明:
 *
 *
 * 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 *
 *
 * 示例:
 *
 * 输入: [3,1,5,8]
 * 输出: 167
 * 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 *
 *
 */

// @lc code=start
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	extendedNums := append([]int{1}, nums...)
	extendedNums = append(extendedNums, 1)

	// dp[i][j]数组表示在 i,j的开区间的最大值
	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, len(nums)+2)
	}

	for length := 1; length <= len(nums); length++ {
		for i := 0; i <= len(nums)-length; i++ {
			j := i + length + 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = maxOf(dp[i][j], extendedNums[k]*extendedNums[i]*extendedNums[j]+dp[i][k]+dp[k][j])
			}
		}
	}

	return dp[0][len(nums)+1]
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
