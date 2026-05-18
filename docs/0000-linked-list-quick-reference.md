# 链表题一页速查版

适用场景：面试前快速过一遍链表高频套路，优先记「怎么判题型、怎么下手、哪些坑最容易炸」。

## 0. 先问三句

1. 我要不要改 `next` 指向？  
2. 我要找的是「位置关系」还是「值关系」？  
3. 退出循环时，指针应该停在 `nil`、尾节点，还是目标前驱？

## 1. 题型到套路（最短映射）

- **遍历统计**：`cur != nil` 直走到底。
- **反转链表**：三指针 `prev cur nxt`。
- **中点/判环/环入口**：快慢指针。
- **倒数第 k / 两链对齐**：双指针同速，先造步数差。
- **删除与局部重连**：优先加 `dummy`，用 `prev` + `cur`。
- **多链合并**：`dummy + tail`，每轮接最优头。

## 2. 模板片段（Go）

### 2.1 全链遍历

```go
for cur := head; cur != nil; cur = cur.Next {
    // use cur
}
```

### 2.2 反转

```go
var prev *ListNode
cur := head
for cur != nil {
    nxt := cur.Next
    cur.Next = prev
    prev = cur
    cur = nxt
}
```

### 2.3 快慢指针常用循环

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
```

### 2.4 删除节点首选骨架

```go
dummy := &ListNode{Next: head}
prev, cur := dummy, head
for cur != nil {
    if needDelete(cur) {
        prev.Next = cur.Next
    } else {
        prev = cur
    }
    cur = cur.Next
}
return dummy.Next
```

## 3. 快慢指针起点速记

- **不是固定必须 `slow=head, fast=head.Next`。**
- 默认先用 `slow=head, fast=head`，因为语义最统一，很多题直接可用。
- 若题目要求前驱位置、左中点、成对比较起点，才改成 `fast=head.Next` 或调整循环条件。

口诀：**先定“slow 要停哪”，再反推 fast 起点和循环条件。**

## 4. 高频易错 8 条

1. 先改 `cur.Next` 再用旧 `cur.Next`，导致断链。  
2. 删除时把 `prev` 误前进，跳过连续待删节点。  
3. 局部反转后忘接回前后两端。  
4. 快慢条件和取值不配套，`nil.Next` 崩溃。  
5. 应处理尾节点却写成 `cur.Next != nil`。  
6. 两链对齐题误用 `cur.Next == nil` 作为切换条件。  
7. 合并题漏写 `tail = tail.Next`。  
8. 分割题忘断尾，残留旧边形成隐式环。  

## 5. 最小自测用例

- `[]`
- `[1]`
- `[1,2]`
- `[1,2,3]`
- `[1,2,2,1]`（回文类）
- 删除头节点场景
- 删除连续节点场景

## 6. 两分钟复习顺序

1. 题型映射看一遍。  
2. 三个模板过一遍（遍历、反转、快慢）。  
3. 高频易错 8 条扫一遍。  
4. 真写前先说清「循环不变量 + 退出态」。  
