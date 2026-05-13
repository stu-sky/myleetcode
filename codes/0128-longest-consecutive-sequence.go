package main

// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/

// longestConsecutive 哈希集合 + 只从序列起点扩展
// 时间 O(n)，空间 O(n)
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, x := range nums {
		set[x] = struct{}{}
	}
	maxLen := 0
	for x := range set {
		if _, ok := set[x-1]; ok {
			continue // x 不是序列起点，跳过
		}
		cur := x
		for {
			if _, ok := set[cur+1]; !ok {
				break
			}
			cur++
		}
		if cur-x+1 > maxLen {
			maxLen = cur - x + 1
		}
	}
	return maxLen
}
