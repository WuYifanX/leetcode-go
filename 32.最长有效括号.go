package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (33.38%)
 * Likes:    883
 * Dislikes: 0
 * Total Accepted:    91.7K
 * Total Submissions: 274.5K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 *
 * 示例 1:
 *
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 *
 *
 * 示例 2:
 *
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	stack := list.New()
	stack.PushBack(-1)
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if stack.Len() != 0 {
				element := stack.Back()
				stack.Remove(element)

				if stack.Len() == 0 {
					stack.PushBack(i)
				} else {
					result = maxOf(result, i-stack.Back().Value.(int))
				}
			}
		} else if s[i] == '(' {
			stack.PushBack(i)
		}
	}
	return result
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

func main() {
	fmt.Println(longestValidParentheses("()(()"))
	// fmt.Println(longestValidParentheses(")()())"))
}
