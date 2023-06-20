/*
这种堆排序，不再每次都将元素添加到尾部，然后上浮翻转，而是在混乱堆的基础上，从底部向上逐层进行下沉操作，下沉操作比较的次数会减少。

步骤如下：
先对最底部的所有非叶子节点进行下沉，即这些非叶子节点与它们的儿子节点比较，较大的儿子和父亲交换位置。
接着从次二层开始的非叶子节点重复这个操作，直到到达根节点最大堆就构建好了。
从底部开始，向上推进，所以这种堆排序又叫自底向上的堆排序。
*/
// TODO: 看懂这代码

package main

import "fmt"

// 先自底向上构建最大堆，再移除堆元素实现堆排序
func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子结点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下移
	end := count - 1

	// 从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start-- // 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
	}

	// 下沉结束了，现在要来排序了
	// 元素大于 2 个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start], 参数 count 只要用来判断是否到堆底，使得下沉结束
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := root*2 + 1

	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续往下沉
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

	HeapSort(list)

	// 打印排序后的值
	fmt.Println(list)
}
