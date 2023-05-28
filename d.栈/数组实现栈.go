// 支持随机访问，查询速度快，但存在元素在数组空间中大量移动的操作，增删效率低。
/*
1.这个增删操作的不就是尾部嘛，用切片感觉效率很高呀
2.说是用数组实际是用切片，使用切片无需指定长度，自动扩容的特性带来了很多方便
*/
package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

// 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size-1]

	// 方式1：切片收缩，但是切片占用的容量不变，所以占用的内存不会减少
	stack.array = stack.array[0 : stack.size-1]

	// 方式2：创建新的数组，新数组的容量就是元素个数,但是时间复杂度增加
	// newArray := make([]string, stack.size-1, stack.size-1)
	// for i := 0; i < stack.size-1; i++ {
	// 	newArray[i] = stack.array[i]
	// }
	// stack.array = newArray

	stack.size = stack.size - 1
	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	v := stack.array[stack.size-1]
	return v
}

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
