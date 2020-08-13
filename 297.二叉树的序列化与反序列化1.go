import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (51.48%)
 * Likes:    331
 * Dislikes: 0
 * Total Accepted:    44.7K
 * Total Submissions: 86.7K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 示例:
 *
 * 你可以将以下二叉树：
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 *
 * 序列化为 "[1,2,3,null,null,4,5]"
 *
 * 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 * 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	maps := make(map[int64]int)

	maps[1] = root.Val
	innerSerialize(root.Right, &maps, 3)
	innerSerialize(root.Left, &maps, 2)
	fmt.Println(maps)

	return stringfy(&maps)
}

func stringfy(maps *map[int64]int) string {
	var build strings.Builder

	for key, value := range *maps {
		build.WriteString(fmt.Sprintf(",%d:%d", key, value))
	}

	return build.String()[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data == "[]" {
		return nil
	}

	array := strings.Split(data, ",")
	maps := make(map[int64]int)

	for _, v := range array {
		keyValueArray := strings.Split(v, ":")
		index, _ := strconv.ParseInt(keyValueArray[0], 10, 64)
		value, _ := strconv.Atoi(keyValueArray[1])
		maps[index] = value
	}

	root := TreeNode{
		Val:   maps[1],
		Right: innerDeserialize(&maps, 3),
		Left:  innerDeserialize(&maps, 2),
	}

	return &root
}

func innerDeserialize(maps *map[int64]int, index int64) *TreeNode {
	value, isOk := (*maps)[index]
	fmt.Println(index, value, isOk)

	if !isOk {
		return nil
	}

	return &TreeNode{
		Val:   value,
		Right: innerDeserialize(maps, index*2+1),
		Left:  innerDeserialize(maps, index*2),
	}
}

func innerSerialize(node *TreeNode, maps *map[int64]int, index int64) {
	if node == nil {
		return
	}
	(*maps)[index] = node.Val
	innerSerialize(node.Left, maps, 2*index)
	innerSerialize(node.Right, maps, 2*index+1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
			Right: &TreeNode{
				Val:   5,
				Right: nil,
				Left:  nil,
			},
		},
	}
	obj := Constructor()
	data := obj.serialize(&root)
	fmt.Println(data)
}
