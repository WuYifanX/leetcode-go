
/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (74.46%)
 * Likes:    338
 * Dislikes: 0
 * Total Accepted:    72.3K
 * Total Submissions: 96.9K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	if k == 0 {
		return result
	}

	isUsed := make([]bool, n+1)
	current := make([]int, 0)

	innerCombine(1, n, k, &current, &isUsed, &result)
	return result
}

func innerCombine(start, end int, k int, current *[]int, isUsed *[]bool, result *[][]int) {

	if len(*current) == k {
		resultArray := make([]int, k)
		copy(resultArray, *current)
		(*result) = append(*result, resultArray)
		return
	}

	for index := start; index <= end; index++ {
		if (*isUsed)[index] {
			continue
		} else {
			(*isUsed)[index] = true
			*current = append((*current), index)
			innerCombine(index+1, end, k, current, isUsed, result)
			(*isUsed)[index] = false
			*current = (*current)[:len(*current)-1]
		}
	}

}

// @lc code=end
