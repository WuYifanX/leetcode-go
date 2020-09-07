/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.01%)
 * Likes:    1283
 * Dislikes: 0
 * Total Accepted:    174.3K
 * Total Submissions: 228.8K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */

// @lc code=start
var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	if n <= 0 {
		return result
	}
	current := make([]byte, 0)
	dfs(n, n, &current)
	return result
}

func dfs(left, right int, current *[]byte) {
	if left == 0 && right == 0 {
		result = append(result, string(*current))
		return
	}

	if left > 0 {
		(*current) = append((*current), '(')
		dfs(left-1, right, current)
		*current = (*current)[:len(*current)-1]
	}

	if left < right {
		(*current) = append((*current), ')')
		dfs(left, right-1, current)
		*current = (*current)[:len(*current)-1]
	}

}

// @lc code=end
