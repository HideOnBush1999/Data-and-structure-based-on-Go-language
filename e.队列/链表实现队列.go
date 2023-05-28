package main

import (
	"fmt"
	"sync"
)

type LinkNode struct {
	Next  *LinkNode
	Value string
}

type LinkQueue struct {
	root *LinkNode // 链表起点
	size int
	lock sync.Mutex
}

// 入队
func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		newNode := new(LinkNode)
		newNode.Value = v

		// 找到链表最后一个节点
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// 尾插法
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	topNode := queue.root
	v := topNode.Value
	// 改变链表起点
	queue.root = topNode.Next
	queue.size = queue.size - 1
	return v
}

// 获取队列的第n个元素
func (queue *LinkQueue) Get(n int) string {
	if queue.size < n {
		panic("Out of range")
	}
	nowNode := queue.root
	for i := 0; i < n; i++ {
		nowNode = nowNode.Next
	}
	return nowNode.Value
}

// 队的大小
func (queue *LinkQueue) Size() int {
	return queue.size
}

// 队是否为空
func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}

func main() {
	queue := new(LinkQueue)
	queue.Add("a")
	queue.Add("b")
	queue.Add("c")
	queue.Add("d")
	fmt.Println(queue.Size())
	fmt.Println(queue.Get(0))
	fmt.Println(queue.Get(1))
	fmt.Println(queue.Remove())
	fmt.Println(queue.Remove())
	fmt.Println(queue.Remove())
	fmt.Println(queue.Size())
}
