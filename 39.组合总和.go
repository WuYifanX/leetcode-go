package main

import "sort"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (69.51%)
 * Likes:    848
 * Dislikes: 0
 * Total Accepted:    129.7K
 * Total Submissions: 186.3K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1：
 *
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2：
 *
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 *
 *
 */

// @lc code=start

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	if len(candidates) == 0 || target < 0 {
		return result
	}

	sort.Sort(IntSlice(candidates))
	current := make([]int, 0)

	backtrack(&candidates, &current, 0, &result, target)

	return result
}

func backtrack(candidates *[]int, current *[]int, start int, result *[][]int, target int) {
	if target == 0 {
		temp := make([]int, len(*current))
		copy(temp, *current)
		*result = append(*result, temp)
		return
	}

	if start == len(*candidates) {
		return
	}

	for index := start; index < len(*candidates); index++ {
		currentValue := (*candidates)[index]
		if len(*current) > 0 && currentValue < (*current)[len(*current)-1] {
			continue
		}

		if currentValue > target {
			return
		}

		*current = append(*current, currentValue)
		// backtrack(candidates, current, start, result, target-currentValue)
		backtrack(candidates, current, start, result, target-currentValue)

		*current = (*current)[:len(*current)-1]
	}

}

// @lc code=end
