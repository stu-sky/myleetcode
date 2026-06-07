package main

// 148. 排序链表
// https://leetcode.cn/problems/sort-list/

// ListNode 由力扣环境提供；本仓库可复用 codes/0160-intersection-of-two-linked-lists.go 中定义。

// sortList 使用自底向上归并排序：按步长 1、2、4... 分段并归并，直到覆盖整表
// 时间 O(n log n)，额外空间 O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}

	dummy := &ListNode{Next: head}

	for step := 1; step < length; step <<= 1 {
		prev := dummy
		curr := dummy.Next

		for curr != nil {
			left := curr
			right := split(left, step)
			curr = split(right, step)

			mergedHead, mergedTail := mergeTwoSortedLists(left, right)
			prev.Next = mergedHead
			prev = mergedTail
		}
	}

	return dummy.Next
}

// split 从 head 开始切出长度为 n 的前半段，返回后半段头结点
// 会在前半段尾结点处断开 Next
func split(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	for i := 1; i < n && head.Next != nil; i++ {
		head = head.Next
	}

	second := head.Next
	head.Next = nil
	return second
}

// mergeTwoSortedLists 归并两个有序链表，返回归并后的头和尾
func mergeTwoSortedLists(l1, l2 *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	for tail.Next != nil {
		tail = tail.Next
	}

	return dummy.Next, tail
}
