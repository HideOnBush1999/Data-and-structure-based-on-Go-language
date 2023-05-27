package main

// 这是模仿可变数组的实现，但是又用到了切片。。。
import (
	"fmt"
	"sync"
)

type Array struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来替代
	len   int        // 真正长度
	cap   int        // 容量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// 初始化数组
func Make(len, cap int) *Array {
	s := new(Array) // new 创建的对象都是指针类型
	if len > cap {
		panic("len large than cap")
	}

	// 把切片当数组用
	array := make([]int, cap, cap)

	// 元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// 添加元素
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 如果需要扩容
	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}

		// 创建新数组，将老数组的数据移动到新数组
		newArray := make([]int, newCap, newCap)

		for k, v := range a.array {
			newArray[k] = v
		}

		a.array = newArray
		a.cap = newCap
	}

	a.array[a.len] = element
	a.len = a.len + 1
}

// 添加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// 获取指定下标元素
func (a *Array) Get(index int) int {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// 返回真实长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() int {
	return a.cap
}

// Print 辅助打印
func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		// 第一个元素
		if i == 0 {
			result = fmt.Sprintf("%s%d", result, array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return // 因为在声明的时候，已经给出了返回变量名，只要return就可以了
}

func main() {
	// 创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
