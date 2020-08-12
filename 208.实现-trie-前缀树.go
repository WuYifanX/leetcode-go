/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (68.10%)
 * Likes:    366
 * Dislikes: 0
 * Total Accepted:    48.4K
 * Total Submissions: 70.9K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
*/

// @lc code=start
type Trie struct {
	cursor []Trie
	isEnd  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{
		cursor: makeCursor(),
		isEnd:  false,
	}
	return trie
}

func makeCursor() []Trie {
	cursor := make([]Trie, 26)

	for i := 0; i < 26; i++ {
		cursor[i] = Trie{
			isEnd:  false,
			cursor: nil,
		}
	}
	return cursor
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	firstLetter := word[0]

	if this.cursor == nil {
		this.cursor = makeCursor()
	}

	if len(word) == 1 {
		this.cursor[firstLetter-'a'].isEnd = true
		return
	}

	this.cursor[firstLetter-'a'].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstLetter := word[0]

	if this.cursor == nil {
		return false
	}

	if len(word) == 1 {
		return this.cursor[firstLetter-'a'].isEnd == true
	}

	return this.cursor[firstLetter-'a'].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	firstLetter := prefix[0]

	if this.cursor == nil {
		return false
	}
	if len(prefix) == 1 {
		currentCursor := this.cursor[firstLetter-'a']
		// fmt.Println("StartsWith", prefix, currentCursor.cursor)
		return currentCursor.cursor != nil || currentCursor.isEnd == true
	}

	return this.cursor[firstLetter-'a'].StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

// func main() {
// 	obj := Constructor()
// 	obj.Insert("apple")
// 	// result := obj.Search("apple")
// 	result := obj.StartsWith("appl")
// 	fmt.Println(result)
// }
