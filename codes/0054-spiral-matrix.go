package main

// 54. 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/

// spiralOrder 按层收缩边界：上 → 右 → 下 → 左，单行/单列用 left>right / top>bottom 防重复
// 时间 O(mn)，除输出外额外空间 O(1)
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1

	out := make([]int, 0, m*n)
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			out = append(out, matrix[top][j])
		}
		top++
		for i := top; i <= bottom; i++ {
			out = append(out, matrix[i][right])
		}
		right--
		if top > bottom {
			break
		}
		if left > right {
			break
		}
		for j := right; j >= left; j-- {
			out = append(out, matrix[bottom][j])
		}
		bottom--
		if left > right {
			break
		}
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			out = append(out, matrix[i][left])
		}
		left++
	}
	return out
}
