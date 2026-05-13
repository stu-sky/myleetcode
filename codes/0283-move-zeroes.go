package main

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/

// moveZeroes 双指针：left 指向下一个非零应放的位置，right 遍历
// 时间 O(n)，空间 O(1)
func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
