package main

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (51.26%)
 * Likes:    445
 * Dislikes: 0
 * Total Accepted:    65.4K
 * Total Submissions: 127.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 *
 *
 *
 * 示例 :
 * 给定二叉树
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 *
 *
 *
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxOf(innerPathCount(root)) - 1
}

// return (maxCountWithout, countWith)
func innerPathCount(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftWithFullPath, leftWithHalfPath := innerPathCount(root.Left)
	rightWithFullPath, rightWithHalfPath := innerPathCount(root.Right)

	fullPath := maxOf(maxOf(leftWithFullPath, rightWithFullPath), 1+leftWithHalfPath+rightWithHalfPath)

	halfPath := maxOf(leftWithHalfPath, rightWithHalfPath) + 1
	return fullPath, halfPath
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
