package main

// 73. 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes/

// setZeroes 第一行 / 第一列作标记区，另用两个布尔量保存「原先首行首列是否需要全清零」
// 时间 O(mn)，空间 O(1)（除若干标量与输入矩阵外）
func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])

	firstRow := false
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRow = true
			break
		}
	}
	firstCol := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
