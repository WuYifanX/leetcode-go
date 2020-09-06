import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (49.09%)
 * Likes:    371
 * Dislikes: 0
 * Total Accepted:    59.5K
 * Total Submissions: 116.2K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 * 说明：
 *
 *
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 *
 *
 * 示例 1:
 *
 * 输入: n = 3, k = 3
 * 输出: "213"
 *
 *
 * 示例 2:
 *
 * 输入: n = 4, k = 9
 * 输出: "2314"
 *
 *
 */

// @lc code=start
var left int
var result string

func getPermutation(n int, k int) string {
	if n <= 0 {
		return ""
	}
	result = ""
	left = k
	isVisit := make([]bool, n+1)
	backtrack(n, &isVisit, "")
	return result
}

func backtrack(n int, isVisit *[]bool, current string) {
	// fmt.Println(current, left)
	if len(result) != 0 {
		return
	}

	if len(current) == n {
		// fmt.Println(current)
		left--
		if left == 0 {
			result = current
			return
		}
	}

	if left > calculateCombination(n-len(current)) {
		// fmt.Println(left, calculateCombination(n-len(current)))
		left -= calculateCombination(n - len(current))
		return
	}

	for i := 1; i <= n; i++ {
		if (*isVisit)[i] == true {
			continue
		}

		(*isVisit)[i] = true
		current += strconv.Itoa(i)
		backtrack(n, isVisit, current)
		current = current[:len(current)-1]
		(*isVisit)[i] = false
	}
}

func calculateCombination(num int) int {
	sum := 1
	for i := num; i > 0; i-- {
		sum *= i
	}
	return sum
}

// @lc code=end

func main() {
	fmt.Println(getPermutation(4, 9))
}
