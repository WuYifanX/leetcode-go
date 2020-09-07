package main

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 *
 * https://leetcode-cn.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (48.97%)
 * Likes:    161
 * Dislikes: 0
 * Total Accepted:    22.9K
 * Total Submissions: 46.6K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 *
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 *
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 *
 * 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 *
 *
 * 示例 2:
 *
 *
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：
 * 把最后一位反向旋转一次即可 "0000" -> "0009"。
 *
 *
 * 示例 3:
 *
 *
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：
 * 无法旋转到目标数字且不被锁定。
 *
 *
 * 示例 4:
 *
 *
 * 输入: deadends = ["0000"], target = "8888"
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 死亡列表 deadends 的长度范围为 [1, 500]。
 * 目标数字 target 不会在 deadends 之中。
 * 每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
 *
 *
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	result := 0
	if target == "0000" {
		return result
	}

	isVisit := make(map[string]bool)
	for _, deadend := range deadends {
		isVisit[deadend] = true
	}

	if isVisit[target] == true || isVisit["0000"] == true {
		return -1
	}

	queue := list.New()
	queue.PushBack("0000")
	var size int
	var current string
	for queue.Len() != 0 {
		size = queue.Len()
		result++
		for size != 0 {
			element := queue.Front()
			queue.Remove(element)
			current = element.Value.(string)
			for index := 0; index < 4; index++ {
				for _, nextNumber := range nextNumbers(current[index]) {
					nextString := current[:index] + string(nextNumber) + current[index+1:]
					if nextString == target {
						return result
					}

					if isVisit[nextString] == true {
						continue
					}

					isVisit[nextString] = true
					queue.PushBack(nextString)
				}
			}

			size--
		}
	}

	return -1
}

func nextNumbers(digit byte) []byte {
	switch digit {
	case '0':
		return []byte{'9', '1'}
	case '9':
		return []byte{'8', '0'}
	case '1':
		return []byte{'2', '0'}
	case '2':
		return []byte{'1', '3'}
	case '3':
		return []byte{'2', '4'}
	case '4':
		return []byte{'3', '5'}
	case '5':
		return []byte{'4', '6'}
	case '6':
		return []byte{'5', '7'}
	case '7':
		return []byte{'6', '8'}
	case '8':
		return []byte{'9', '7'}
	}
	return nil
}

// @lc code=end
