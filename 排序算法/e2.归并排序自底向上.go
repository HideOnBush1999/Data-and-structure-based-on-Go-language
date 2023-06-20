package main

import "fmt"

func MergeSort(array []int, begin int, end int) {
	// 步数为1开始，step 长度的数组表示一个有序的数组
	step := 1

	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin; i < end; i += step << 1 {
			low := i                  // 第一个有序数组的上界
			mid := low + step         // 第一个有序数组的下界，第二个有序数组的上界
			high := low + (step << 1) // 第二个有序数组的下界

			// 不存在第二个数组,直接返回
			if mid > end {
				return
			}

			// 第二个数组长度不够
			if high > end {
				high = end
			}
			// 两个有序数组进行合并
			fmt.Println("step:", step, " i:", i, array[low:mid], array[mid:high])
			merge(array, low, mid, high)
		}

		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}
}

func merge(array []int, begin int, mid int, end int) {
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	l, r := 0, 0
	result := make([]int, 0, newSize) // 三个参数对应切片类型，初始长度，容量

	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边元素
		rValue := array[mid+r]   // 右边元素
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

	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14}
	fmt.Println(list)
	MergeSort(list, 0, len(list))
	fmt.Println(list)
}

// 自底向上非递归排序，我们可以看到没有递归那样程序栈的增加，效率比自顶向上的递归版本高
