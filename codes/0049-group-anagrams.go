package main

// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/

// groupAnagrams 字符计数法：用每个字符串的字符频次作为分组 key
// 时间 O(n*k)，空间 O(n*k)
func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		mp[count] = append(mp[count], s)
	}
	result := make([][]string, 0, len(mp))
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}
