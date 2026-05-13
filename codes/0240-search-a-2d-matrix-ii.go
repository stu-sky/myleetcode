package main

// 240. 搜索二维矩阵 II
// https://leetcode.cn/problems/search-a-2d-matrix-ii/

// searchMatrix 从「左下角」出发：偏大则上移一行，偏小则右移一列（对称地可从右上角开始）
// 每步排除一行或一列，至多 m+n 步 → 时间 O(m+n)，额外空间 O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	i, j := m-1, 0
	for i >= 0 && j < n {
		v := matrix[i][j]
		if v == target {
			return true
		}
		if v > target {
			i--
		} else {
			j++
		}
	}
	return false
}
