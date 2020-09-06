package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (54.31%)
 * Likes:    903
 * Dislikes: 0
 * Total Accepted:    172.6K
 * Total Submissions: 311.4K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
var result []string
var dictMap map[byte][]byte

func letterCombinations(digits string) []string {
	result = make([]string, 0)

	if len(digits) == 0 {
		return result
	}

	dictMap = make(map[byte][]byte)
	dictMap['2'] = []byte{'a', 'b', 'c'}
	dictMap['3'] = []byte{'d', 'e', 'f'}
	dictMap['4'] = []byte{'g', 'h', 'i'}
	dictMap['5'] = []byte{'j', 'k', 'l'}
	dictMap['6'] = []byte{'m', 'n', 'o'}
	dictMap['7'] = []byte{'p', 'q', 'r', 's'}
	dictMap['8'] = []byte{'t', 'u', 'v'}
	dictMap['9'] = []byte{'w', 'x', 'y', 'z'}

	current := ""
	dfs(digits, current)
	return result
}

func dfs(digits string, current string) {
	if len(current) == len(digits) {
		result = append(result, current)
		return
	}

	for _, digit := range dictMap[digits[len(current)]] {
		current = current + string(digit)
		dfs(digits, current)
		current = current[:len(current)-1]
	}
}

// @lc code=end
