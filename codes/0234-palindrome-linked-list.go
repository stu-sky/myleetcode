package main

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/

// ListNode 由力扣环境提供；与 codes/0160-intersection-of-two-linked-lists.go 中定义一致时可合并编译。

// isPalindrome：快慢指针定「前半尾结点」，断开思路下反转后半，双指针逐项比；最后再反转一次恢复原链
// 时间 O(n)，额外空间 O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfHead := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfHead
	ok := true
	for ok && p2 != nil {
		if p1.Val != p2.Val {
			ok = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfHead)
	return ok
}

// endOfFirstHalf 快慢步长二比一；偶链结束时 slow 为左半截尾结点；奇链时 slow 为中点结点
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}
