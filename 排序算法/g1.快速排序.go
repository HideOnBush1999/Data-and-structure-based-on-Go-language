/*
	快速排序通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

步骤如下：
先从数列中取出一个数作为基准数。一般取第一个数。
分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
再对左右区间重复第二步，直到各区间只有一个数。

切片优化：
切分的结果极大地影响快速排序的性能，比如每次切分的时候选择的基数数都是数组中最大或者最小的，会出现最坏情况，为了避免切分不均匀情况的发生，有几种方法改进：
随机基准数选择：每次进行快速排序切分时，先将数列随机打乱，再进行切分，这样随机加了个震荡，减少不均匀的情况。当然，也可以随机选择一个基准数，而不是选第一个数。
中位数选择：每次取数列头部，中部，尾部三个数，取三个数的中位数为基准数进行切分。
*/
package main

import (
	"fmt"
)

func QuickSort(array []int, begin, end int) {
	if begin < end { // 退出递归条件
		// 当数组小于 4 时使用直接插入排序
		if end-begin <= 4 {
			fmt.Println("进入直接插入排序：", array[begin:end+1])
			InsertSort(array[begin : end+1])
			fmt.Println("直接插入排序后：", array[begin:end+1])
			return
		}

		// 进行切分
		fmt.Println("切分前： ", array[begin:end+1]) // 因为左闭右开，要看到最后一个元素需要 +1
		loc := partition(array, begin, end)
		fmt.Println("切分后： ", array[begin:end+1], "loc: ", loc)
		// 对左部分进行快排
		QuickSort(array, begin, loc-1)
		// 对右部分进行快排
		QuickSort(array, loc+1, end)
	}
}

/*
5 [9] 1 6 8 14 6 49 25 4 6 [3]  因为 9 > 5，两个[]交换位置后，右边[]左移
5 [3] 1 6 8 14 6 49 25 4 [6] 9  因为 3 !> 5，两个[]不需要交换，左边[]右移
5 3 [1] 6 8 14 6 49 25 4 [6] 9  因为 1 !> 5，两个[]不需要交换，左边[]右移
5 3 1 [6] 8 14 6 49 25 4 [6] 9  因为 6 > 5，两个[]交换位置后，右边[]左移
5 3 1 [6] 8 14 6 49 25 [4] 6 9  因为 6 > 5，两个[]交换位置后，右边[]左移
5 3 1 [4] 8 14 6 49 [25] 6 6 9  因为 4 !> 5，两个[]不需要交换，左边[]右移
5 3 1 4 [8] 14 6 49 [25] 6 6 9  因为 8 > 5，两个[]交换位置后，右边[]左移
5 3 1 4 [25] 14 6 [49] 8 6 6 9  因为 25 > 5，两个[]交换位置后，右边[]左移
5 3 1 4 [49] 14 [6] 25 8 6 6 9  因为 49 > 5，两个[]交换位置后，右边[]左移
5 3 1 4 [6] [14] 49 25 8 6 6 9  因为 6 > 5，两个[]交换位置后，右边[]左移
5 3 1 4 [14] 6 49 25 8 6 6 9  两个[]已经汇总，因为 14 > 5，所以 5 和[]之前的数 4 交换位置
第一轮切分结果：4 3 1 5 14 6 49 25 8 6 6 9
*/
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

func InsertSort(list []int) {
	n := len(list)

	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 一直往左边找，比排序大的数都往后挪，腾空位给待排序插入
		for ; j >= 0 && deal < list[j]; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = deal // 结束了，待排序的数插入空位
	}
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

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}
