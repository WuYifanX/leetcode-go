/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.76%)
 * Likes:    734
 * Dislikes: 0
 * Total Accepted:    127K
 * Total Submissions: 163K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, make([]int, 0))

	if len(nums) == 0 {
		return result
	}

	for length := 1; length <= len(nums); length++ {
		current := make([]int, 0)
		innerCombine(&nums, length, 0, &current, &result)
	}

	return result
}

func innerCombine(nums *[]int, k, start int, current *[]int, result *[][]int) {

	if len(*current) == k {
		resultArray := make([]int, k)
		copy(resultArray, *current)
		(*result) = append(*result, resultArray)
		return
	}

	for index := start; index < len(*nums); index++ {

		*current = append((*current), (*nums)[index])
		innerCombine(nums, k, index+1, current, result)
		*current = (*current)[:len(*current)-1]
	}

}

// @lc code=end
