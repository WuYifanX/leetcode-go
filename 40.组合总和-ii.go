import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (62.49%)
 * Likes:    349
 * Dislikes: 0
 * Total Accepted:    79.8K
 * Total Submissions: 127.6K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

// @lc code=start

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	if len(candidates) == 0 || target < 0 {
		return result
	}

	sort.Sort(IntSlice(candidates))
	isUsed := make([]bool, len(candidates))

	current := make([]int, 0)

	backtrack(&candidates, &current, &isUsed, 0, &result, target)

	return result
}

func backtrack(candidates *[]int, current *[]int, isUsed *[]bool, start int, result *[][]int, target int) {
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
		// already sorted, so if currentValue is larger than target, it will never meet the requirement
		if currentValue > target {
			return
		}

		// is Used, so continue
		if (*isUsed)[index] == true {
			continue
		}

		// if index -1 and index are the same number && index -1 is not be used, so the index -1 and index should be the same
		if index > 0 && (*candidates)[index] == (*candidates)[index-1] && (*isUsed)[index-1] == false {
			continue
		}

		(*isUsed)[index] = true
		*current = append(*current, currentValue)

		backtrack(candidates, current, isUsed, index+1, result, target-currentValue)

		(*isUsed)[index] = false
		*current = (*current)[:len(*current)-1]
	}

}

// @lc code=end
