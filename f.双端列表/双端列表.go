package main

import (
	"fmt"
	"sync"
)

// 链表节点 （和双向链表、循环链表是一样的）
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string
}

// 双端列表
type DoubleList struct {
	head *ListNode // 指向链表头部
	tail *ListNode // 指向链表尾部
	len  int
	lock sync.Mutex
}

// 链表的第一个元素的前驱节点为 nil，最后一个元素的后驱节点也为 nil

// 链表节点的方法
// -------------------------------------------------------------------
// GetValue 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// GetPre 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// GetNext 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// HashNext 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// HashPre 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.pre != nil
}

// IsNil 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

// 双端链表的方法
// ---------------------------------------------------------------
// Len 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

// 从头部开始某个位置前插入新节点
// AddNodeFromHead 从头部开始，添加节点到第N+1个元素之前，
// N=0表示添加到第一个元素之前，表示新节点成为新的头部，
// N=1表示添加到第二个元素之前，以此类推 （简单点，N都是大于等于0的）
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超过或等于列表长度，一定找不到 n=0是特殊情况
	if n != 0 && n >= list.len {
		panic("index out")
	}

	// 找到第 n+1 个位置的元素
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 新节点
	newNode := new(ListNode)
	newNode.value = v

	// 1.如果定位到的节点为空，表示列表为空
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		pre := node.pre

		// 2.如果定位到的节点前驱为nil，那么定位到的节点为链表头部
		if pre.IsNil() {
			newNode.next = node
			node.pre = newNode
			list.head = newNode
		} else {
			// 3.如果定位到的节点不是链表头部
			pre.next = newNode
			newNode.pre = pre

			newNode.next = node
			node.pre = newNode
		}
	}
	list.len = list.len + 1
}

// 从尾部开始某个位置前插入新节点
// 添加节点到第N+1个元素之后，N=0表示添加到第一个元素之后，表示新节点成为新的尾部
// N=1表示添加到第二个元素之后
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n != 0 && n >= list.len {
		panic("index out")
	}

	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	newNode := new(ListNode)
	newNode.value = v

	// 1.链表为空
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		next := node.next

		// 2.链表不为空，定位到的节点为尾结点
		if next.IsNil() {
			node.next = newNode
			newNode.pre = node

			list.tail = newNode
		} else {
			// 3.链表不为空，定位到的节点不是尾结点
			node.next = newNode
			newNode.pre = node

			newNode.next = next
			next.pre = newNode
		}
	}
	list.len = list.len + 1
}

// 返回头节点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// 返回尾结点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}

// 从头部开始往后找，获取第N+1个位置的节点，索引从0开始
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	if n >= list.len {
		panic("index out")
	}

	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}
	return node
}

// 从头部开始往后找，获取第N+1个位置的节点，索引从0开始
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	if n >= list.len {
		panic("index out")
	}

	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	return node
}

// PopFromHead 从头部开始往后找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len {
		return nil
	}

	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	pre := node.pre
	next := node.next

	// 1.既是头节点又是尾结点，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 2.是头节点
		next.pre = nil
		list.head = next
	} else if next.IsNil() {
		// 3.是尾结点
		pre.next = nil
		list.tail = pre
	} else {
		// 4.是中间节点
		pre.next = next
		next.pre = pre
	}
	list.len = list.len - 1
	return node
}

// PopFromTail 从头部开始往后找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromTail(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n >= list.len {
		return nil
	}

	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	pre := node.pre
	next := node.next

	// 1.既是头节点又是尾结点，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		// 2.是头节点
		next.pre = nil
		list.head = next
	} else if next.IsNil() {
		// 3.是尾结点
		pre.next = nil
		list.tail = pre
	} else {
		// 4.是中间节点
		pre.next = next
		next.pre = pre
	}
	list.len = list.len - 1
	return node
}

func main() {
	list := new(DoubleList)
	// 在列表头部插入新元素
	list.AddNodeFromHead(0, "I")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")
	list.AddNodeFromTail(list.Len()-1, "begin second")
	list.AddNodeFromHead(list.Len()-1, "end second")
	// 正常遍历，比较慢，因为内部会遍历拿到值返回
	for i := 0; i < list.Len(); i++ {
		// 从头部开始索引
		node := list.IndexFromHead(i)
		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsNil() {
			fmt.Println(node.GetValue())
		}
	}
	fmt.Println("----------")
	// 正常遍历，特别快，因为直接拿到的链表节点
	// 先取出第一个元素
	first := list.First()
	for !first.IsNil() {
		// 如果非空就一直遍历
		fmt.Println(first.GetValue())
		// 接着下一个节点
		first = first.GetNext()
	}
	fmt.Println("----------")
	// 元素一个个 POP 出来
	for {
		node := list.PopFromHead(0)
		if node.IsNil() {
			// 没有元素了，直接返回
			break
		}
		fmt.Println(node.GetValue())
	}
	fmt.Println("----------")
	fmt.Println("len", list.Len())
}
