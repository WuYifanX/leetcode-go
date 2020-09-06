
/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (42.31%)
 * Likes:    544
 * Dislikes: 0
 * Total Accepted:    83.1K
 * Total Submissions: 195.8K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true
 * 给定 word = "SEE", 返回 true
 * 给定 word = "ABCB", 返回 false
 *
 *
 *
 * 提示：
 *
 *
 * board 和 word 中只包含大写和小写英文字母。
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

// @lc code=start
var directions [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var result bool

func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}

	result = false
	isVisited := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		isVisited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == word[0] {
				isVisited[i][j] = true
				dfs(&grid, i, j, &isVisited, &word, 1)
				isVisited[i][j] = false
			}
		}
	}
	return result
}

func dfs(grid *[][]byte, x, y int, isVisited *[][]bool, word *string, index int) {
	if index == len(*word) || result == true {
		result = true
		return
	}

	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if isValidPosition(newX, newY, len(*grid), len((*grid)[0])) {
			if (*isVisited)[newX][newY] {
				continue
			}
			if (*grid)[newX][newY] == (*word)[index] {
				(*isVisited)[newX][newY] = true
				dfs(grid, newX, newY, isVisited, word, index+1)
				(*isVisited)[newX][newY] = false
			}
		}
	}
}
func isValidPosition(x, y, maxX, maxY int) bool {
	return x >= 0 && y >= 0 && x < maxX && y < maxY
}

// @lc code=end
