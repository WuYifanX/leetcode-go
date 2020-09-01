/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (76.67%)
 * Likes:    862
 * Dislikes: 0
 * Total Accepted:    181.7K
 * Total Submissions: 236.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 0 {
		return result
	}
	// value == 1 is used; value == 0, is not used
	isUsed := make([]int, len(nums))
	current := make([]int, 0)
	innerPermute(&nums, 0, &isUsed, &current, &result)
	return result
}

func innerPermute(nums *[]int, currentIndex int, isUsed *[]int, current *[]int, result *[][]int) {

	if currentIndex == len(*nums) {
		newArray := make([]int, len(*nums))
		copy(newArray, *current)

		(*result) = append(*result, newArray)
		return
	}

	for i := 0; i < len(*nums); i++ {
		if (*isUsed)[i] == 1 {
			continue
		}

		(*isUsed)[i] = 1
		(*current) = append((*current), (*nums)[i])
		innerPermute(nums, currentIndex+1, isUsed, current, result)
		(*isUsed)[i] = 0
		(*current) = (*current)[:len((*current))-1]
	}
}

// @lc code=end
