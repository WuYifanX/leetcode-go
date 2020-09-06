/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (42.18%)
 * Likes:    362
 * Dislikes: 0
 * Total Accepted:    70.6K
 * Total Submissions: 167.1K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
 *
 * 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
 *
 * 示例:
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * 运行你的函数后，矩阵变为：
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * 解释:
 *
 * 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 */

// @lc code=start
var directions [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func solve(grid [][]byte) {
	if len(grid) == 0 {
		return
	}

	isVisited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		isVisited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				if isVisited[i][j] == false && grid[i][j] == 'O' {
					dfs(&grid, i, j, &isVisited)
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'Y' {
				grid[i][j] = 'O'
			} else if grid[i][j] == 'O' {
				grid[i][j] = 'X'
			}
		}
	}
}

func dfs(grid *[][]byte, x, y int, isVisited *[][]bool) {
	(*isVisited)[x][y] = true

	(*grid)[x][y] = 'Y'

	for _, direction := range directions {
		newX, newY := x+direction[0], y+direction[1]
		if isValidPosition(newX, newY, len(*grid), len((*grid)[0])) {
			if (*isVisited)[newX][newY] {
				continue
			}
			if (*grid)[newX][newY] == 'O' {
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
	solve([][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}})
}
