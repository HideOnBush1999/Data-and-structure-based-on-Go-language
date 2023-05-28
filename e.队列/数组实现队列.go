package main

import (
	"fmt"
	"sync"
	"unsafe"
)

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

// 入队
func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)
	queue.size = queue.size + 1
}

// 出队
func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}
	v := queue.array[0]

	// 方式1
	// queue.array = queue.array[1:queue.size]

	// 方式2
	// for i := 1; i < queue.size; i++ {
	// 	queue.array[i-1] = queue.array[i]
	// }
	// queue.array = queue.array[0 : queue.size-1]
	// 方式3
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	queue.size = queue.size - 1
	return v
}

// 获取队列的第n个元素
func (queue *ArrayQueue) Get(n int) string {
	return queue.array[n]
}

// 队的大小
func (queue *ArrayQueue) Size() int {
	return queue.size
}

// 队是否为空
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

func main() {
	queue := ArrayQueue{
		array: []string{"a", "b", "c"},
		size:  3,
	}
	fmt.Printf("切片内存大小： %v\n", unsafe.Sizeof(cap(queue.array)))
	fmt.Println(queue.Get(2))
	// 尝试添加元素
	queue.Add("d")
	queue.Add("e")
	queue.Add("f")
	queue.Add("g")
	queue.Add("h")
	queue.Add("i")
	queue.Add("j")
	fmt.Printf("切片内存大小： %v\n", unsafe.Sizeof(cap(queue.array)))
	fmt.Println("队列的长度为： ", queue.size)
	fmt.Println(queue.Remove()) // 使用方式一进行出队
	fmt.Println(queue.Remove()) // 使用方式一进行出队
	fmt.Printf("切片内存大小： %v\n", unsafe.Sizeof(cap(queue.array)))
}

/*
Q1: 计算方式1实现出队，切片内存的变化
答： 切片收缩，但是切片占用的容量不变，所以占用的内存不会减少
Q2: 为什么queue是ArrayQueue类型，而不是该类型的指针也可以使用其方法
在Go语言中，当你将一个结构体类型的变量作为接收者(receiver)来定义方法时，Go编译器会自动为该类型的指针和值两种类型生成方法集。这意味着，无论你使用结构体类型的指针还是值作为接收者，你都可以调用方法。

在你的代码中，queue是ArrayQueue类型的变量，而不是指针。然而，由于Go语言的自动转换机制，你仍然可以使用该变量调用结构体的方法。

当你调用queue.Remove()时，编译器会隐式地将queue转换为对应的指针类型&queue，以便在Remove()方法中正确地操作结构体字段。这种转换在调用方法时是透明的，你不需要显式地使用指针。

所以，尽管Remove()方法的接收者是quere *ArrayQueue，你仍然可以使用queue.Remove()来调用该方法，而不需要显式地取地址。

这种隐式转换使得代码更加简洁和易读。你可以根据个人的喜好和需求，选择将接收者定义为指针类型还是值类型。如果你想在方法内部修改结构体字段的值，可以使用指针类型的接收者；如果不需要修改结构体字段，可以使用值类型的接收者，这样你可以使用结构体的变量直接调用方法。

*/
