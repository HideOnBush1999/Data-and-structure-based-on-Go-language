package main

import (
	"fmt"
)

func QuickSort(array []int, begin, end int) {
	if begin < end { // 退出递归条件
		// 进行切分
		loc := partition(array, begin, end)
		// 那边元素少先排那边
		if loc-begin < end-loc {
			// 先排左边
			QuickSort(array, begin, loc-1)
			begin = loc + 1 // 下一次循环排序的范围从切分点的右侧开始，因为左侧的元素已经在当前循环中被排序好了
		} else {
			// 先排右边
			QuickSort(array, loc+1, end)
			end = loc - 1 // 下一次循环排序的范围从切分点的左侧开始，因为右侧的元素已经在当前循环中被排序好了
		}
	}
}

// 切分函数，并返回切分元素的下标
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
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

/* TODO:  这个真的有用嘛？
先排元素少的那边，最后不是还要排元素多的那边，这和不判断直接先排左边，再排右边有什么区别？
*/
