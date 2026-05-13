package main

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// lengthOfLongestSubstring 标准滑动窗口模板
// 时间 O(n)，空间 O(字符集大小)
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int) // 窗口内每个字符的出现次数
	left := 0
	ans := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		for window[c] > 1 { // 窗口内有重复字符，收缩左边界
			window[s[left]]--
			left++
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
