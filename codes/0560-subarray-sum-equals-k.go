package main

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/

// subarraySum 前缀和 + 哈希表
// cnt[t] = 遍历到当前元素之前，前缀和等于 t 出现过几次
// 子数组和为 k ⇔ 当前前缀和 sum 与左边界前前缀和 prev 满足 sum - prev = k，即查 cnt[sum-k]
// 时间 O(n)，空间 O(n)
func subarraySum(nums []int, k int) int {
	cnt := make(map[int]int)
	cnt[0] = 1 // 空前缀和为 0，算出现 1 次（子数组从下标 0 开始时用）
	sum := 0
	ans := 0
	for _, x := range nums {
		sum += x
		ans += cnt[sum-k] // 前面有多少个位置，使得「那一段结尾到当前」的和正好是 k
		cnt[sum]++
	}
	return ans
}
