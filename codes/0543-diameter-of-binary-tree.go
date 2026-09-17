package main

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/
// TreeNode 由 LeetCode 环境提供。

// diameterOfBinaryTree 后序遍历：返回单链深度，更新双侧路径长度。
// 时间 O(n)，额外空间 O(h)，h 为树高。
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0 // 已检查过的节点中，最长路径的边数；不是 depth 的返回值。
	// depth 返回从当前节点向下的最长单链节点数，并顺便更新外部答案。
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		// 职责一：检查经过当前节点的路径。节点数为 left+right+1，边数为 left+right。
		if left+right > ans {
			ans = left + right
		}
		// 职责二：返回当前节点的深度。向父节点延伸时只能选一边，再加当前节点。
		if left > right {
			return left + 1
		}
		return right + 1
	}
	depth(root) // 根的深度不是直径；调用的目的是遍历所有节点并更新 ans。
	return ans
}
