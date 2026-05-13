package main

// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/

// minWindow 滑动窗口：右扩直到覆盖 t，再左缩求最短合法窗口
// 时间 O(|s|+|t|)，空间 O(|字符集|)，本题字母可视为常数
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	needTypes := len(need) // t 中有几种字符要「达标」

	window := make(map[byte]int)
	valid := 0 // 几种字符在窗口内已经达到 need 中的个数

	left, right := 0, 0
	start, length := 0, len(s)+1 // 哨兵：无合法窗口时 length 会大于 len(s)

	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == needTypes {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length > len(s) {
		return ""
	}
	return s[start : start+length]
}
