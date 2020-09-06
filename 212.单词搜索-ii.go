
import "fmt"

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (41.95%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    18.4K
 * Total Submissions: 43.8K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 * 示例:
 *
 * 输入:
 * words = ["oath","pea","eat","rain"] and board =
 * [
 * ⁠ ['o','a','a','n'],
 * ⁠ ['e','t','a','e'],
 * ⁠ ['i','h','k','r'],
 * ⁠ ['i','f','l','v']
 * ]
 *
 * 输出: ["eat","oath"]
 *
 * 说明:
 * 你可以假设所有输入都由小写字母 a-z 组成。
 *
 * 提示:
 *
 *
 * 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
 * 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
 * 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
 *
 *
*/

// @lc code=start

type Trie struct {
	isEnd bool
	next  [](*Trie)
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	trie := Trie{}
	for _, word := range words {
		buildTrie(word, 0, &trie)
	}
	result := make(map[string]bool)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			isVisit := make(map[string]bool)
			isVisit[transferPositionToString([]int{i, j})] = true
			dfs(&board, &isVisit, i, j, &trie, &result, "")
		}
	}

	resultArray := make([]string, 0)

	for key := range result {
		resultArray = append(resultArray, key)
	}
	return resultArray
}

func isValidPosition(position []int, m, n int) bool {
	return position[0] >= 0 && position[0] < m && position[1] >= 0 && position[1] < n
}

func dfs(board *[][]byte, isVisit *map[string]bool, x, y int, trie *Trie, result *map[string]bool, currentWord string) {
	// fmt.Println(isVisit, x, y, currentWord)
	// fmt.Println()
	currentChar := (*board)[x][y]
	nextTrie := (*trie).next[currentChar-'a']
	// fmt.Println(nextTrie)
	if nextTrie == nil {
		return
	}

	if nextTrie.isEnd {
		(*result)[currentWord+string(currentChar)] = true
	}

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n := len(*board), len((*board)[0])

	for _, direction := range directions {
		nextPosition := []int{x + direction[0], y + direction[1]}
		if isValidPosition(nextPosition, m, n) && !(*isVisit)[transferPositionToString(nextPosition)] {
			// fmt.Println((*isVisit)[transferPositionToString(nextPosition)], nextPosition)
			(*isVisit)[transferPositionToString(nextPosition)] = true
			dfs(board, isVisit, nextPosition[0], nextPosition[1], nextTrie, result, currentWord+string(currentChar))
			(*isVisit)[transferPositionToString(nextPosition)] = false

		}
	}
}

func transferPositionToString(position []int) string {
	return fmt.Sprintf("%d,%d", position[0], position[1])
}

func buildTrie(word string, index int, trie *Trie) {

	if index == len(word) {
		trie.isEnd = true
		return
	}

	current := word[index]

	if trie.next == nil {
		trie.next = make([](*Trie), 26)
	}

	if trie.next[current-'a'] == nil {
		trie.next[current-'a'] = &Trie{
			isEnd: false,
			next:  make([](*Trie), 26),
		}
	}

	buildTrie(word, index+1, trie.next[current-'a'])
}

// @lc code=end

// func main() {
// 	fmt.Println(findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"acdb"}))
// }
