package main

// 138. 随机链表的复制
// https://leetcode.cn/problems/copy-list-with-random-pointer/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 使用「穿插复制」三遍扫描：先复制并穿插，再补 random，最后拆分新旧链表
// 时间 O(n)，额外空间 O(1)（不计输出链表）
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一遍：每个原节点后插入一个复制节点
	for curr := head; curr != nil; {
		copyNode := &Node{Val: curr.Val, Next: curr.Next}
		curr.Next = copyNode
		curr = copyNode.Next
	}

	// 第二遍：设置复制节点的 Random
	for curr := head; curr != nil; curr = curr.Next.Next {
		copyNode := curr.Next
		if curr.Random != nil {
			copyNode.Random = curr.Random.Next
		}
	}

	// 第三遍：拆分为原链表和复制链表
	newHead := head.Next
	for curr := head; curr != nil; {
		copyNode := curr.Next
		curr.Next = copyNode.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
		curr = curr.Next
	}

	return newHead
}
