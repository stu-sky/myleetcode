package main

import "sort"

// 15. 三数之和
// https://leetcode.cn/problems/3sum/

// threeSum 排序 + 固定一个数 + 对撞双指针
// 时间 O(n²)，空间 O(log n)（排序栈空间）
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break // 最小值已 > 0，不可能凑出 0
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复的第一个数
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // 跳过重复的第二个数
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // 跳过重复的第三个数
				}
				left++
				right--
			}
		}
	}
	return result
}
