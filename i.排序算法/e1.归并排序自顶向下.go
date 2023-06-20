/*
通过递归地先使每个子序列有序，再将有序的序列进行合并成一个有序的序列。
*/
package main

import "fmt"

// 自顶向下归并排序，排序范围在 [begin, end)
func MergeSort(array []int, begin int, end int) {
	// 元素数量小于1时才进入递归
	if end-begin > 1 {
		// 将数组一分为二
		// [1, 5)  mid = 1 + 5/2 = 3
		// [1, 6)  mid = 1 + 6/2 = 4
		mid := begin + (end-begin+1)/2

		// 先将左边排序好
		fmt.Println("左边排序：", begin, mid, array[begin:mid])
		MergeSort(array, begin, mid)

		// 再将右边排序好
		fmt.Println("右边排序：", mid, end, array[mid:end])
		MergeSort(array, mid, end)

		// 两个有序数组进行合并
		fmt.Println("开始合并：", begin, mid, end, array[begin:mid], array[mid:end])
		merge(array, begin, mid, end)
	}
}

// 归并操作，将两个有序的子数组 [begin,mid) 和 [mid, end)合并成一个有序的数组
func merge(array []int, begin int, mid int, end int) {
	leftSize := mid - begin         // 左边数组的长度，因为是一开一闭所以不用加1
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else if rValue < lValue {
			result = append(result, rValue)
			r++
		} else {
			result = append(result, lValue)
			result = append(result, rValue)
			l++
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	// l最大值 --> leftSize 	  begin+l --> begin+leftSize --> mid
	// r最大值 --> rightSize      mid+r --> mid+rightSize --> end
	if l == leftSize { // 左边先被遍历完
		result = append(result, array[mid+r:end]...) // 这个数组使用本身也是左闭右开
	} else {
		result = append(result, array[begin+l:mid]...)
	}

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	// 在Go语言中，不需要手动释放切片内存，因为Go的垃圾回收器会自动管理内存的分配和释放。
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

func main() {
	// list := []int{5}
	// MergeSort(list, 0, len(list))
	// fmt.Println(list)

	list1 := []int{9, 5, 1, 2, 3, 56, 6}
	MergeSort(list1, 0, len(list1))
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, 0, len(list2))
	fmt.Println(list2)
}
