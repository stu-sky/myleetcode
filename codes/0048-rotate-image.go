package main

// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image/

// rotate 原地顺时针转 90°：先按主对角线转置，再把每一行左右翻转
// 时间 O(n^2)，额外空间 O(1)（沿用题面「使用额外内存」为常数的常见口径）
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
