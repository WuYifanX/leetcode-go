import (
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
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	array := strings.Split(data, ",")

	root, _ := innerDeserialize(&array, 0)
	return root
}

func innerDeserialize(array *[]string, index int) (*TreeNode, int) {
	if index >= len(*array) || (*array)[index] == "#" {
		return nil, 1
	}

	value, _ := strconv.Atoi((*array)[index])
	leftRoot, leftCount := innerDeserialize(array, index+1)
	rightRoot, rightCount := innerDeserialize(array, index+1+leftCount)

	return &TreeNode{
		Val:   value,
		Left:  leftRoot,
		Right: rightRoot,
	}, leftCount + rightCount + 1
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

// func main() {
// 	root := TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val:   2,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 		Right: &TreeNode{
// 			Val: 3,
// 			Left: &TreeNode{
// 				Val:   4,
// 				Right: nil,
// 				Left:  nil,
// 			},
// 			Right: &TreeNode{
// 				Val:   5,
// 				Right: nil,
// 				Left:  nil,
// 			},
// 		},
// 	}
// 	obj := Constructor()
// 	data := obj.serialize(&root)
// 	result := obj.deserialize(data)
// 	fmt.Println(data)
// 	fmt.Println(result)
// }
