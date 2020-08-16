import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 *
 * https://leetcode-cn.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (66.33%)
 * Likes:    435
 * Dislikes: 0
 * Total Accepted:    38K
 * Total Submissions: 57.2K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * 假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。
 * 编写一个算法来重建这个队列。
 *
 * 注意：
 * 总人数少于1100人。
 *
 * 示例
 *
 *
 * 输入:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 *
 * 输出:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 *
 *
 */

// @lc code=start

type IntSlice [][]int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool {
	if s[i][0] != s[j][0] {
		return s[i][0]-s[j][0] > 0
	} else {
		return s[i][1]-s[j][1] < 0
	}
}

func reconstructQueue(people [][]int) [][]int {
	peopleIntSlice := IntSlice(people)
	result := make([][]int, 0)
	sort.Sort(peopleIntSlice)

	for i := 0; i < len(peopleIntSlice); i++ {
		result = SliceInsert(result, peopleIntSlice[i][1], peopleIntSlice[i])
	}
	return result
}

func SliceInsert(s [][]int, index int, value []int) [][]int {
	rear := append([][]int{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

// @lc code=end
