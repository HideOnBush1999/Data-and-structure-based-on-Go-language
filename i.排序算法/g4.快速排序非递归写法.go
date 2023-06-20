/*
非递归写法仅仅是将之前的递归栈转化为自己维持的手工栈。
(更加体会到了递归是以栈的形式存储的含义)
*/

package main

import (
	"fmt"
	"sync"
)

// 链表栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 锁
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value int
}

// 入栈(将新节点放在栈顶)
func (stack *LinkStack) Push(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否者新元素插入链表的头部
		// 原来的链表
		preNode := stack.root

		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 原来的链表链接到新的元素后面
		newNode.Next = preNode

		// 将新节点放在头部
		stack.root = newNode
	}
	stack.size++
}

// 出栈(将栈顶元素弹出)
func (stack *LinkStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 顶部元素要出栈
	topNode := stack.root
	v := topNode.Value

	// 将顶部元素的后续链接链上
	stack.root = topNode.Next

	// 栈中元素数量 -1
	stack.size = stack.size - 1
	return v
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LinkStack) Print() {
	NowNode := stack.root
	for NowNode != nil {
		fmt.Printf("%d  ", NowNode.Value)
		NowNode = NowNode.Next
	}
	fmt.Printf("\n")
}

// 一直往栈中存入左边和右边的元素，然后弹出作为切分函数的参数，在切分函数中，数组的排列顺序发生变化
// 当栈中元素为空，说明全部排序结束，退出
func QuickSort(array []int) {
	// 人工栈
	helpStack := new(LinkStack)

	// 第一次初始化栈，推入下标0，len(array-1),表示第一次对全数组范围切分
	helpStack.Push(len(array) - 1)
	helpStack.Push(0)
	helpStack.Print()

	// 栈非空证明存在未排序的部分
	for !helpStack.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := helpStack.Pop() // 范围区间左边
		end := helpStack.Pop()   // 范围区间右边

		// 进行切分, 将基准值放在下标为 loc 的位置上
		loc := partition(array, begin, end) // 在 partition 函数的内部，array的排列顺序发生了变化
		// fmt.Printf("begin: %d, loc: %d, end: %d\n", begin, loc, end)

		// 右边范围入栈
		if loc+1 < end { // 当基准值右边的元素个数大于2时，右边的元素还需要排序
			helpStack.Push(end)     // 先放入右边的下标，再放入左边的下标
			helpStack.Push(loc + 1) // 因为弹出的时候我们先得到左边下标再得到右边下标
		}
		// fmt.Printf("右边范围入栈后：")
		// helpStack.Print()

		// 左边范围入栈
		if begin < loc-1 { // 当基准值左边的元素个数大于2时，左边的元素还需要排序
			helpStack.Push(loc - 1)
			helpStack.Push(begin)
		}
		// fmt.Printf("左边范围入栈后：")
		// helpStack.Print()
	}
}

// 切分函数
func partition(array []int, begin, end int) int {
	i := begin + 1 // 取名叫左下标
	j := end       // 取名叫右下标

	// 没重合之前
	for i < j {
		if array[i] > array[begin] { // 如果左下标元素比基准值大，那么左下标元素与右下标元素互换，右下标左移
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else { // 如果左下标元素比基准值小，那么左下标右移
			i++
		}
	}
	/* 跳出 for 循环后， i = j,此时数组被分割成两个部分
	-----> array[begin+1] ~ array[i-1]  < array[begin]
	-----> array[i+1] ~ array[end] > array[begin]
	这个时候将数组 array 分成两个部分，再将 array[i] 与 array[begin] 进行比较，决定 array[i] 的位置
	最后将 array[i] 与 array[begin]交换，进行两个分割部分的排序，直到最后 i = j 不满足条件就退出
	*/
	if array[i] >= array[begin] {
		// 当前值大于等于基准值的话，与基准值的前一个元素交换
		i--
	}

	// 将基准值放在比其小和比其大的中间
	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list) // 在 QuickSort() 中的 partition()， list 是引用传递，被改变
	fmt.Println(list)
}
