/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (59.62%)
 * Likes:    385
 * Dislikes: 0
 * Total Accepted:    84.8K
 * Total Submissions: 141.9K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}

	isUsed := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count, isExist := isUsed[nums[i]]
		if isExist {
			isUsed[nums[i]] = count + 1
		} else {
			isUsed[nums[i]] = 1
		}
	}
	current := make([]int, 0)
	innerPermuteUnique(&nums, 0, &current, &isUsed, &result)

	return result
}

func innerPermuteUnique(nums *[]int, currentIndex int, current *[]int, isUsed *map[int]int, result *[][]int) {
	if currentIndex == len(*nums) {
		newArray := make([]int, len(*nums))
		copy(newArray, *current)
		*result = append(*result, newArray)
	}

	for key, count := range *isUsed {
		if count == 0 {
			continue
		} else {
			(*isUsed)[key] = count - 1
			(*current) = append(*current, key)
			innerPermuteUnique(nums, currentIndex+1, current, isUsed, result)
			(*isUsed)[key] = count
			(*current) = (*current)[:len(*current)-1]
		}
	}
}

// @lc code=end
