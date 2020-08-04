package main

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}

	dict := make([]int, 26)

	for i := 0; i < len(s); i++ {
		dict[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		dict[t[i]-'a'] -= 1

		if dict[t[i]-'a'] == -1 {
			return t[i]
		}
	}
	return t[0]
}

// @lc code=end
