// package main

import "fmt"

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 *
 * https://leetcode-cn.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (31.13%)
 * Likes:    487
 * Dislikes: 0
 * Total Accepted:    49.3K
 * Total Submissions: 158.4K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
 *
 * '?' 可以匹配任何单个字符。
 * '*' 可以匹配任意字符串（包括空字符串）。
 *
 *
 * 两个字符串完全匹配才算匹配成功。
 *
 * 说明:
 *
 *
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
 *
 *
 * 示例 1:
 *
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 *
 * 示例 2:
 *
 * 输入:
 * s = "aa"
 * p = "*"
 * 输出: true
 * 解释: '*' 可以匹配任意字符串。
 *
 *
 * 示例 3:
 *
 * 输入:
 * s = "cb"
 * p = "?a"
 * 输出: false
 * 解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
 *
 *
 * 示例 4:
 *
 * 输入:
 * s = "adceb"
 * p = "*a*b"
 * 输出: true
 * 解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
 *
 *
 * 示例 5:
 *
 * 输入:
 * s = "acdcb"
 * p = "a*c?b"
 * 输出: false
 *
 */

// @lc code=start
func isMatch(s string, p string) bool {
	if len(p) == 0 && len(s) == 0 {
		return true
	}

	if len(p) == 0 {
		return false
	}

	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {

			if i == 0 || j == 0 {
				if j == 0 {
					dp[0][0] = true
				} else if p[j-1] == '*' {
					dp[i][j] = dp[i][j-1]
				}
				continue
			}

			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end
func main() {
	// isMatch("a", "a*")
	fmt.Println(isMatch("ho", "**ho"))
}
