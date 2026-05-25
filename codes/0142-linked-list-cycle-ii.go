package main

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/

// ListNode 由力扣环境提供；本仓库可复用 codes/0160-intersection-of-two-linked-lists.go 中定义。

// detectCycle 先用快慢指针判是否有环；若相遇，将一指针移回头结点，再同速前进，相遇点即环入口
// 时间 O(n)，额外空间 O(1)
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
