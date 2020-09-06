import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (49.19%)
 * Likes:    405
 * Dislikes: 0
 * Total Accepted:    78.6K
 * Total Submissions: 158.5K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
 *
 * 有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是
 * "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 * 输入：s = "1111"
 * 输出：["1.1.1.1"]
 *
 *
 * 示例 4：
 *
 * 输入：s = "010010"
 * 输出：["0.10.0.10","0.100.1.0"]
 *
 *
 * 示例 5：
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 3000
 * s 仅由数字组成
 *
 *
 */

// @lc code=start
var result []string
var position []int

func restoreIpAddresses(s string) []string {
	result = make([]string, 0)

	if len(s) < 4 || len(s) > 12 {
		return result
	}

	position = make([]int, 0)
	backtrack(s, 0)
	return result
}

func backtrack(s string, lastPosition int) {
	if len(position) == 3 {
		if isValidIPAddress(s[position[2]:]) {
			result = append(result, buildIP(s, position))
		}
	}

	for i := lastPosition + 1; i < len(s) && i <= lastPosition+3; i++ {
		if isValidIPAddress(s[lastPosition:i]) {
			position = append(position, i)
			backtrack(s, i)
			position = position[:len(position)-1]
		}
	}
}

func buildIP(s string, position []int) string {
	var build strings.Builder

	build.WriteString(s[0:position[0]])
	build.WriteString(".")
	for i := 1; i < len(position); i++ {
		build.WriteString(s[position[i-1]:position[i]])
		build.WriteString(".")
	}
	build.WriteString(s[position[2]:])
	return build.String()
}

func isValidIPAddress(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	if s[0] == '0' {
		return false
	}

	value, err := strconv.Atoi(s)

	if err != nil || value > 255 {
		return false
	}

	return true
}

// @lc code=end

// func main() {
// 	fmt.Println(restoreIpAddresses("25525511135"))
// }
