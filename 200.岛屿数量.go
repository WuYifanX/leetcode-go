import "fmt"

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (50.07%)
 * Likes:    745
 * Dislikes: 0
 * Total Accepted:    153.2K
 * Total Submissions: 305K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ['1','1','1','1','0'],
 * ['1','1','0','1','0'],
 * ['1','1','0','0','0'],
 * ['0','0','0','0','0']
 * ]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ['1','1','0','0','0'],
 * ['1','1','0','0','0'],
 * ['0','0','1','0','0'],
 * ['0','0','0','1','1']
 * ]
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */

// @lc code=start
var directions [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func numIslands(grid [][]byte) int {
	result := 0
	if len(grid) == 0 {
		return result
	}

	isVisited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		isVisited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isVisited[i][j] == false && grid[i][j] == '1' {
				dfs(&grid, i, j, &isVisited)
				result++
			}
		}
	}

	return result
}

func dfs(grid *[][]byte, x, y int, isVisited *[][]bool) {
	(*isVisited)[x][y] = true

	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if isValidPosition(newX, newY, len(*grid), len((*grid)[0])) {
			if (*isVisited)[newX][newY] {
				continue
			}
			if (*grid)[x][y] == '1' {
				dfs(grid, newX, newY, isVisited)
			}
		}
	}
}
func isValidPosition(x, y, maxX, maxY int) bool {
	return x >= 0 && y >= 0 && x < maxX && y < maxY
}

// @lc code=end
func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}}))
}
