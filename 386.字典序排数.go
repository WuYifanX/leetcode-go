package main

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {

	result := make([]int, 0)

	if n <= 9 {
		for i := 1; i <= n; i++ {
			result = append(result, i)
		}
		return result
	}

	for i := 1; i <= 9; i++ {
		dfsForLexicalOrder(i, n, &result)
	}
	return result
}

func dfsForLexicalOrder(current, max int, result *[]int) {
	if current > max {
		return
	}
	(*result) = append(*result, current)
	for i := 0; i < 10; i++ {
		nextRoot := current*10 + i

		if nextRoot > max {
			break
		}
		dfsForLexicalOrder(nextRoot, max, result)
	}
}

// @lc code=end
