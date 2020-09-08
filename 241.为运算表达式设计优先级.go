// package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(input string) []int {
	result := make([]int, 0)

	if len(input) == 0 {
		return make([]int, 0)
	}
	hasOperation := false

	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' || input[i] == '/' {
			hasOperation = true
			lefts := diffWaysToCompute(input[:i])
			rights := diffWaysToCompute(input[i+1:])
			for _, left := range lefts {
				for _, right := range rights {
					switch input[i] {
					case '+':
						{
							result = append(result, left+right)
						}
					case '-':
						{
							result = append(result, left-right)
						}
					case '*':
						{
							result = append(result, left*right)
						}
					case '/':
						{
							result = append(result, left/right)
						}
					}
				}
			}
		}
	}

	if hasOperation == false {
		intValue, _ := strconv.Atoi(input)
		result = append(result, intValue)
	}
	return result
}

// @lc code=end
