package main

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/

// maxSlidingWindow 单调双端队列（存下标）：队首到队尾对应 nums 值单调递减
// 时间 O(n)，空间 O(k)
func maxSlidingWindow(nums []int, k int) []int {
	var dq []int // 下标队列，队头为当前窗口最大值下标
	ans := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 队头已滑出窗口
		for len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}
		// 新元素更大或相等时，队尾更小的不可能再成为窗口最大值
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)

		if i >= k-1 {
			ans = append(ans, nums[dq[0]])
		}
	}
	return ans
}
