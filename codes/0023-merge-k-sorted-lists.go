package main

// 23. 合并 K 个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/

// ListNode 由力扣环境提供；本仓库可复用 codes/0160-intersection-of-two-linked-lists.go 中定义。

// mergeKLists 使用分治两两合并：按间隔 1、2、4... 把 lists 中链表成对合并
// 时间 O(N log k)，额外空间 O(1)（不计返回链表；迭代写法无递归栈）
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}

	for interval := 1; interval < n; interval *= 2 {
		for i := 0; i+interval < n; i += interval * 2 {
			lists[i] = mergeTwoLists23(lists[i], lists[i+interval])
		}
	}

	return lists[0]
}

func mergeTwoLists23(l1, l2 *ListNode) *ListNode {
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

	return dummy.Next
}
