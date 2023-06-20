// 只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快
package main

import (
	"fmt"
	"sync"
)

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 链表栈
type LinkStack struct {
	root *LinkNode // 链表起点
	size int
	lock sync.Mutex
}

// 入栈(将新节点放在栈顶)
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		preNode := stack.root

		newNode := new(LinkNode)
		newNode.Value = v

		// 头插法
		newNode.Next = preNode
		// 更新头部
		stack.root = newNode
	}
	stack.size = stack.size + 1
}

// 出栈(将栈顶元素弹出)
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("stack is empty")
	}

	topNode := stack.root
	v := topNode.Value
	// 更新头部
	stack.root = topNode.Next
	stack.size = stack.size - 1
	return v
}

// 获取栈顶元素
func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}
	return stack.root.Value
}

// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}
