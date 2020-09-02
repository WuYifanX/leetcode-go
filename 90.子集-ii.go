import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (60.70%)
 * Likes:    296
 * Dislikes: 0
 * Total Accepted:    46.8K
 * Total Submissions: 77.2K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 */

// @lc code=start

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, make([]int, 0))

	if len(nums) == 0 {
		return result
	}

	sort.Sort(IntSlice(nums))

	for length := 1; length <= len(nums); length++ {
		current := make([]int, 0)
		isUsed := make([]bool, len(nums))
		innerCombine(&nums, length, 0, &current, &isUsed, &result)
	}

	return result

}

func innerCombine(nums *[]int, k, start int, current *[]int, isUsed *[]bool, result *[][]int) {

	if len(*current) == k {
		resultArray := make([]int, k)
		copy(resultArray, *current)
		(*result) = append(*result, resultArray)
		return
	}

	for index := start; index < len(*nums); index++ {
		if index > 0 && (*nums)[index] == (*nums)[index-1] && (*isUsed)[index-1] == false {
			continue
		}
		(*isUsed)[index] = true
		*current = append((*current), (*nums)[index])
		innerCombine(nums, k, index+1, current, isUsed, result)
		*current = (*current)[:len(*current)-1]
		(*isUsed)[index] = false
	}

}

// @lc code=end
