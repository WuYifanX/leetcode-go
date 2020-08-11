
/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (59.69%)
 * Likes:    1028
 * Dislikes: 0
 * Total Accepted:    74K
 * Total Submissions: 123.9K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	} else if n == 0 {
		return m
	}

	// i, j表示word1[i-1], word2[j-1]的最小编辑距离
	dp := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			replaceOperationTimes := 1
			if i == 0 {
				dp[0][j] = j
				continue
			} else if j == 0 {
				dp[i][0] = i
				continue
			}

			if word1[i-1] == word2[j-1] {
				replaceOperationTimes = 0
			}

			replaceOperations := replaceOperationTimes + dp[i-1][j-1]
			deleteOperations := 1 + dp[i-1][j]
			insertOperations := minOf(1+dp[i][j-1], replaceOperationTimes+dp[i-1][j-1])

			dp[i][j] = minOf(replaceOperations, minOf(deleteOperations, insertOperations))
		}
	}
	return dp[m][n]
}

func minOf(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end

// func main() {
// 	fmt.Println(minDistance("horse", "ros"))
// 	fmt.Println(minDistance("intention", "execution"))
// 	fmt.Println(minDistance("plasma", "altruism"))
// }
