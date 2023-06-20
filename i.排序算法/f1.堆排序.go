/*
优先队列是一种能完成以下任务的队列：插入一个数值，取出最小或最大的数值（获取数值，并且删除）

优先队列可以用二叉树来实现，我们称这种结构为二叉堆。
最小堆和最大堆是二叉堆的一种，是一棵完全二叉树（一种平衡树）。

最小堆的性质：
父节点的值都小于左右儿子节点。
这是一个递归的性质。

最大堆的性质：
父节点的值都大于左右儿子节点。
这是一个递归的性质。

最大堆实现细节(两个操作)：
push：向堆中插入数据时，首先在堆的末尾插入数据，如果该数据比父亲节点还大，那么交换，然后不断向上提升，直到没有大小颠倒为止。
pop：从堆中删除最大值时，首先把最后一个值复制到根节点上，并且删除最后一个数值，然后和儿子节点比较，如果值小于儿子，与儿子节点交换，然后不断向下交换， 直到没有大小颠倒为止。在向下交换过程中，如果有两个子儿子都大于自己，就选择较大的。
*/

package main

import "fmt"

type Heap struct {
	// 堆的大小
	Size int
	// 使用切片来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

// 初始化一个堆
func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array // 浅拷贝
	return h
}

// 最大堆插入元素
func (h *Heap) Push(x int) {
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	// i 是要插入节点的下标
	i := h.Size

	for i > 0 {
		parent := (i - 1) / 2

		// 如果插入的值小于等于父亲节点，那么可以直接退出循环
		if x <= h.Array[parent] {
			break
		}

		// 如果插入的值大于父亲节点，父亲节点被拉下来
		h.Array[i] = h.Array[parent]
		i = parent
	}
	h.Array[i] = x
	h.Size++
}

// 最大堆移除元素
func (h *Heap) Pop() int {
	// 没有元素，返回-1 (当最大值就是-1的时候，那要怎么办)
	if h.Size == 0 {
		return -1
	}

	// 取出根节点
	ret := h.Array[0]

	// 因为根节点要被删除了，将最后一个节点放在根节点的位置上
	h.Size--
	x := h.Array[h.Size]  // 将最后一个元素的值先拿出来
	h.Array[h.Size] = ret // 将移除的元素放在最后一个元素的位置上(这样就将最大值放在最后面，正序排序)

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	i := 0
	for {
		// a, b 为下标 i 左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 左儿子下标超出了,表示没有左子树，那么右子树也没有，直接返回
		if a >= h.Size {
			break
		}

		// 有右子树，拿到两个子节点中较大节点的下标
		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b // a 表示较大的那个节点
		}

		// 父亲节点的值都大于或等于两个儿子较大的那个，不需要向下继续翻转了，返回
		if x >= h.Array[a] {
			break
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[a]

		// 继续往下操作
		i = a
	}
	// 将最后一个元素的值 x 放在不会再翻转的位置
	h.Array[i] = x
	return ret
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	// 构建最大堆
	h := NewHeap(list)
	for _, v := range list {
		h.Push(v)
	}
	// fmt.Println(h.Array)

	// 将堆元素移除
	for range list {
		// fmt.Println("pop前：", list)
		h.Pop() // 每 Pop 一次都会将最大值（根节点）放在最后（最后节点），并保持最大堆的结构
		// fmt.Println("pop后：", list)
	}

	// 打印排序后的值
	fmt.Println(list)
	// Q: 为什么 list 会被改变？
	// A: 因为 Pop() 中对 *Heap 类型的变量中切片数组进行了改变。
	// 而 h.Array = array 是浅拷贝，所以导致了 list 也被改变了
}
