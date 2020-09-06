import (
	"bytes"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 *
 * https://leetcode-cn.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Medium (65.35%)
 * Likes:    202
 * Dislikes: 0
 * Total Accepted:    24.9K
 * Total Submissions: 38.1K
 * Testcase Example:  '"a1b2"'
 *
 * 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
 *
 *
 *
 * 示例：
 * 输入：S = "a1b2"
 * 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * 输入：S = "3z4"
 * 输出：["3z4", "3Z4"]
 *
 * 输入：S = "12345"
 * 输出：["12345"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * S 的长度不超过12。
 * S 仅由数字和字母组成。
 *
 *
 */

// @lc code=start
var result []string

func letterCasePermutation(S string) []string {
	result = make([]string, 0)
	if len(S) == 0 {
		return result
	}

	byteArray := bytes.ToLower([]byte(S))
	current := ""
	dfs(&byteArray, &current)

	return result
}

func dfs(byteArray *[]byte, current *string) {
	if len((*current)) == len(*byteArray) {
		result = append(result, string(*current))
		return
	}

	newCurrent := (*current) + string((*byteArray)[len((*current))])
	dfs(byteArray, &newCurrent)

	newCurrent = newCurrent[:len(newCurrent)-1]
	if (*byteArray)[len(newCurrent)] >= 'a' && (*byteArray)[len(newCurrent)] <= 'z' {
		newCurrent += string((*byteArray)[len(newCurrent)] - 32)
		dfs(byteArray, &newCurrent)
	}
}

// @lc code=end

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
