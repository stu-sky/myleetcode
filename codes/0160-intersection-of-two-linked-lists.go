package main

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 「A 链走完接 B」「B 链走完接 A」：两指针对齐路径差后在交点相会；无交点则同抵 nil
// 时间 O(m+n)，额外空间 O(1)；交点按节点地址相同界定，不以 Val 相同为准
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
