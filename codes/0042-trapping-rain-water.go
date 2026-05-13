package main

// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/

// --------------- 解法一：动态规划（最好理解） ---------------
// 先预处理每列左边最高和右边最高，再逐列算积水
// 时间 O(n)，空间 O(n)
func trapDP(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max42(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max42(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += min42(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

// --------------- 解法二：双指针 ---------------
// 从两端向中间走，矮的一端水位已确定，直接算
// 时间 O(n)，空间 O(1)
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

// --------------- 解法三：单调栈 ---------------
// 维护单调递减栈，遇到更高柱子时弹栈，按「横向层」计算积水
// 时间 O(n)，空间 O(n)
func trapStack(height []int) int {
	var stack []int
	ans := 0

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break // 左边没有墙，接不住水
			}
			left := stack[len(stack)-1]
			w := i - left - 1
			h := min42(height[left], height[i]) - height[top]
			ans += w * h
		}
		stack = append(stack, i)
	}
	return ans
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min42(a, b int) int {
	if a < b {
		return a
	}
	return b
}
