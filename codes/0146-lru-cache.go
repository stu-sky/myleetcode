package main

// 146. LRU 缓存
// https://leetcode.cn/problems/lru-cache/

type lruNode struct {
	key  int
	val  int
	prev *lruNode
	next *lruNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*lruNode
	head     *lruNode // 最近使用端（伪头）
	tail     *lruNode // 最久未使用端（伪尾）
}

func Constructor(capacity int) LRUCache {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*lruNode, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node := &lruNode{key: key, val: value}
	this.cache[key] = node
	this.addToHead(node)

	if len(this.cache) > this.capacity {
		removed := this.removeTail()
		delete(this.cache, removed.key)
	}
}

func (this *LRUCache) addToHead(node *lruNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *lruNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *lruNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
