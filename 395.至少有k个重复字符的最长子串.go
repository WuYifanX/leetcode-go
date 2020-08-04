package main

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/description/
 *
 * algorithms
 * Medium (43.36%)
 * Likes:    186
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 27K
 * Testcase Example:  '"aaabb"\n3'
 *
 * 找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 * s = "aaabb", k = 3
 *
 * 输出:
 * 3
 *
 * 最长子串为 "aaa" ，其中 'a' 重复了 3 次。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s = "ababbc", k = 2
 *
 * 输出:
 * 5
 *
 * 最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
 *
 *
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	if k <= 1 {
		return len(s)
	}
	if len(s) < k || len(s) == 0 {
		return 0
	}

	dict := make([]int, 26)
	for i := 0; i < len(s); i++ {
		dict[s[i]-'a'] += 1
	}

	result := 0

	splitPoints := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if dict[s[i]-'a'] < k {
			splitPoints = append(splitPoints, i)
		}
	}

	if len(splitPoints) == 0 {
		return len(s)
	}

	left := 0
	for i := 0; i < len(splitPoints); i++ {
		splitPoint := splitPoints[i]

		// fmt.Println(splitPoint, splitPoints, s[left:splitPoint], s)
		result = maxForLongestSubstring(result, longestSubstring(s[left:splitPoint], k))
		left = splitPoint + 1
	}
	result = maxForLongestSubstring(result, longestSubstring(s[left:], k))

	return result
}

func maxForLongestSubstring(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// func main() {
// 	fmt.Println(longestSubstring("bbaaacbd", 3))
// }
