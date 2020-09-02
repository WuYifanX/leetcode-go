
import "fmt"

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode-cn.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (44.37%)
 * Likes:    359
 * Dislikes: 0
 * Total Accepted:    40K
 * Total Submissions: 90.1K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或
 * -中选择一个符号添加在前面。
 *
 * 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
 *
 *
 *
 * 示例：
 *
 * 输入：nums: [1, 1, 1, 1, 1], S: 3
 * 输出：5
 * 解释：
 *
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 *
 * 一共有5种方法让最终目标和为3。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 数组非空，且长度不会超过 20 。
 * 初始的数组的和不会超过 1000 。
 * 保证返回的最终结果能被 32 位整数存下。
 *
 *
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make(map[string]int)

	return backtrack(&nums, 0, S, &memo)
}

func backtrack(nums *[]int, index int, target int, memo *map[string]int) int {
	key := fmt.Sprint("%d,%d", index, target)
	v, isOk := (*memo)[key]
	if isOk {
		return v
	}

	if len(*nums) == index {
		if target == 0 {
			(*memo)[key] = 1
		} else {
			(*memo)[key] = 0
		}
		return (*memo)[key]
	}
	(*memo)[key] = backtrack(nums, index+1, target+(*nums)[index], memo) + backtrack(nums, index+1, target-(*nums)[index], memo)
	return (*memo)[key]
}

// @lc code=end
