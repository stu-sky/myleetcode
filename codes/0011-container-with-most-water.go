package main

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/

// maxArea 对撞双指针：每次移动较矮的一端
// 时间 O(n)，空间 O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		h := min(height[left], height[right])
		area := h * (right - left)
		if area > ans {
			ans = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
