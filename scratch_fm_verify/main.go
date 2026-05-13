package main

import (
	"fmt"
)

// 与仓库 codes/0041-first-missing-positive.go 一致（用于穷尽小用例）

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		x := nums[i]
		if x < 0 {
			x = -x
		}
		if x <= n && nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func brute(nums []int) int {
	for want := 1; ; want++ {
		ok := false
		for _, v := range nums {
			if v == want {
				ok = true
				break
			}
		}
		if !ok {
			return want
		}
	}
}

func clone(xs []int) []int {
	out := make([]int, len(xs))
	copy(out, xs)
	return out
}

func testOne(orig []int) {
	cp := clone(orig)
	got := firstMissingPositive(cp)
	want := brute(orig)
	if got != want {
		panic(fmt.Sprintf("bad: %v got %d want %d after %v mut", orig, got, want, cp))
	}
}

func dfs(out []int, maxLen int, minV int, maxV int) {
	testOne(out)
	if len(out) == maxLen {
		return
	}
	for v := minV; v <= maxV; v++ {
		dfs(append(out, v), maxLen, minV, maxV)
	}
}

func main() {
	maxLen := 5
	minV := -2
	maxV := 8
	fmt.Println("exhaustive small arrays...")
	dfs(make([]int, 0), maxLen, minV, maxV)
	fmt.Println("OK")
}
