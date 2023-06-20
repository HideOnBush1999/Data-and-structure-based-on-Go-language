/*  排序数列可能存在大量重复值，使用三向切分快速排序
将数组分成三部分，大于基准数，等于基准数，小于基准数，这个时候需要维护三个下标

针对重复率高的时候，避免相同元素来回交换，节省交换次数。
*/

package main

import "fmt"

// 三切分的快速排序
func QuickSort(array []int, begin int, end int) {
	if begin < end { // 退出条件
		// 三向切分函数，返回左边和右边下标（两个下标之间是等于基准数的数，原先的切分函数只要返回一个下标就好了）
		lt, gt := partition(array, begin, end)
		// 左边三向快排
		QuickSort(array, begin, lt-1)
		// 右边三向快排
		QuickSort(array, gt+1, end)
	}
}

/*
[4] [8] 2 4 4 4 7 [9]  从中间[]开始：8 > 4，中右[]进行交换，右边[]左移
[4] [9] 2 4 4 4 [7] 8  从中间[]开始：9 > 4，中右[]进行交换，右边[]左移
[4] [7] 2 4 4 [4] 9 8  从中间[]开始：7 > 4，中右[]进行交换，右边[]左移
[4] [4] 2 4 [4] 7 9 8  从中间[]开始：4 == 4，不需要交换，中间[]右移
[4] 4 [2] 4 [4] 7 9 8  从中间[]开始：2 < 4，中左[]需要交换，中间和左边[]右移
2 [4] 4 [4] [4] 7 9 8  从中间[]开始：4 == 4，不需要交换，中间[]右移
2 [4] 4 4 [[4]] 7 9 8  从中间[]开始：4 == 4，不需要交换，中间[]右移，因为已经重叠了
*/
func partition(array []int, begin int, end int) (int, int) {
	// 这个算法中是有三个下标的，返回两个下标
	lt := begin       // 左下标
	gt := end         // 右下标
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数

	// 中间下标小于等于右边下标
	for i <= gt {
		if array[i] > v { // 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v { // 小于基准值，那么交换，左指针和中间指针右移
			array[i], array[lt] = array[lt], array[i]
			i++
			lt++
		} else { // 等于基准值
			i++
		}
		fmt.Println(array[begin : end+1])
	}
	return lt, gt
}

func main() {
	// list := []int{5}
	// QuickSort(list, 0, len(list)-1)
	// fmt.Println(list)

	// list1 := []int{5, 9}
	// QuickSort(list1, 0, len(list1)-1)
	// fmt.Println(list1)

	// list2 := []int{5, 9, 1}
	// QuickSort(list2, 0, len(list2)-1)
	// fmt.Println(list2)

	list3 := []int{5, 9, 1, 5, 8, 14, 5, 49, 25, 5}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}
