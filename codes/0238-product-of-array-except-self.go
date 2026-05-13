package main

// 238. 除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/

// productExceptSelf 两遍扫描：answer[i] = 左侧乘积 × 右侧乘积
// 第一遍写入左前缀积，第二遍用变量累乘右后缀并乘入 answer
// 时间 O(n)，空间 O(1)（除输出数组）
func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	left := 1
	for i := 0; i < n; i++ {
		answer[i] = left
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}
