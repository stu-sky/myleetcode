package main

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

// ListNode 由力扣环境提供；本仓库可复用 codes/0160-intersection-of-two-linked-lists.go 中定义。

// reverseKGroup 每次检查后面是否有 k 个结点，够则原地翻转这一组并接回链表；不足 k 个保持原样
// 时间 O(n)，额外空间 O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	groupPrev := dummy

	for {
		kth := groupPrev
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}
		if kth == nil {
			break
		}

		groupNext := kth.Next
		prev := groupNext
		curr := groupPrev.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		newGroupHead := kth
		newGroupTail := groupPrev.Next
		groupPrev.Next = newGroupHead
		groupPrev = newGroupTail
	}

	return dummy.Next
}
