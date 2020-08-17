/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.81%)
 * Likes:    556
 * Dislikes: 0
 * Total Accepted:    65.6K
 * Total Submissions: 146.3K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 *
 * 示例 1 :
 *
 *
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 *
 *
 * 说明 :
 *
 *
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 *
 *
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 0
	// maps[sum] = count
	maps := make(map[int]int)
	maps[0] = 1

	sum := 0
	// sum start from i....j
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		target := sum - k
		count, isOk := maps[target]

		if isOk {
			result += count
		}

		count, isOk = maps[sum]
		if isOk {
			maps[sum] = count + 1
		} else {
			maps[sum] = 1
		}
	}
	return result
}

// @lc code=end
