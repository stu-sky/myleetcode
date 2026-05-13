package main

// 41. 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive/

// firstMissingPositive 标记法：负数表示「该下标对应的正数出现过」
// 先置无关值为 n+1；再据当前格子的绝对值去给对应下标打勾；最后找第一个仍为正的格子
// 时间 O(n)，空间 O(1)（除输入数组外）
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		x := nums[i]
		if x < 0 {
			x = -x
		}
		if x <= n && nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
