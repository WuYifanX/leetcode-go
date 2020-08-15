
/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 *
 * https://leetcode-cn.com/problems/remove-invalid-parentheses/description/
 *
 * algorithms
 * Hard (47.42%)
 * Likes:    224
 * Dislikes: 0
 * Total Accepted:    12.9K
 * Total Submissions: 27.1K
 * Testcase Example:  '"()())()"'
 *
 * 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
 *
 * 说明: 输入可能包含了除 ( 和 ) 以外的字符。
 *
 * 示例 1:
 *
 * 输入: "()())()"
 * 输出: ["()()()", "(())()"]
 *
 *
 * 示例 2:
 *
 * 输入: "(a)())()"
 * 输出: ["(a)()()", "(a())()"]
 *
 *
 * 示例 3:
 *
 * 输入: ")("
 * 输出: [""]
 *
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {

	if isValid(s) == true {
		return []string{s}
	}

	return innerRemoveInvalidParentheses([]string{s})
}

func innerRemoveInvalidParentheses(strings []string) []string {
	removedMap := make(map[string]bool)
	isAnyValid := false
	for _, s := range strings {
		for i := 0; i < len(s); i++ {
			removedString := ""

			if i == len(s)-1 {
				removedString = s[0:i]
			} else {
				removedString = s[0:i] + s[i+1:]
			}

			isElementValid, isOk := removedMap[removedString]
			// fmt.Println(isValid(removedString), removedString)
			if (isOk && isElementValid) || isValid(removedString) {
				isAnyValid = true
				removedMap[removedString] = true
			} else {
				removedMap[removedString] = false
			}
		}
	}
	// fmt.Println(strings, removedMap)

	result := make([]string, 0)

	if isAnyValid {
		for key, isValid := range removedMap {
			if isValid {
				result = append(result, key)
			}
		}
		return result
	} else {
		for key := range removedMap {
			result = append(result, key)
		}
		return innerRemoveInvalidParentheses(result)
	}
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

// @lc code=end

// func main() {
// 	fmt.Println(removeInvalidParentheses("))"))
// }
