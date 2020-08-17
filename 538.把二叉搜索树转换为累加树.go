
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 *
 * https://leetcode-cn.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (62.87%)
 * Likes:    305
 * Dislikes: 0
 * Total Accepted:    35.2K
 * Total Submissions: 55.9K
 * Testcase Example:  '[5,2,13]'
 *
 * 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater
 * Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
 *
 *
 *
 * 例如：
 *
 * 输入: 原始二叉搜索树:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * 输出: 转换为累加树:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
 *
 *
 *
 * 注意：本题和 1038:
 * https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	innerConvertBST(root, 0)
	return root
}

//        5
//      /   \
//    2     13
//   / \   /  \
//  1   3 10   15

//        5
//      /   \
//    2     28
//   / \   /  \
//  1   3 38   15

//        2
//      /   \
//    0      3
//   / \
//  -4  1

//        5
//      /   \
//    6      3
//   / \
//  2   6

//        2
//      /   \
//    1     6
//   / \    / \
// -2 null nul 4

func innerConvertBST(node *TreeNode, valueFromParent int) int {
	if node == nil {
		return valueFromParent
	}

	rightChildSum := innerConvertBST(node.Right, valueFromParent)
	node.Val += rightChildSum
	leftChildSum := innerConvertBST(node.Left, node.Val)

	return leftChildSum
}

// @lc code=end
