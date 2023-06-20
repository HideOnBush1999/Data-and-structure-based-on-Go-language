/* 两点改进
1. 对于小规模数组，使用直接插入排序
2. 原地排序（通过翻转算法实现），节约掉辅助数组空间的占用
*/

/*
	翻转算法（手摇算法）：1.将前部分逆序 2.将后部分逆序 3.对整体逆序

将字符串 abcde1234567 的前 5 个字符与后面的字符交换位置
1. 分成两部分：[abcde][1234567]
2. 分别逆序变成：[edcba][7654321]
3. 整体逆序：[1234567abcde]
归并原地排序利用了手摇算法的特征，不需要额外的辅助数组。
*/
package main

import "fmt"

// 直接插入排序
func InsertSort(list []int) {
	n := len(list)

	// 进行 N -1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 一直往左边找，比排序大的数都往后挪。腾空位给待排序插入
		for ; j >= 0 && deal < list[j]; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = deal
	}
}

// 自底向上归并算法优化版 TODO:
func MergeSort(array []int, n int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(array[a:b])
		a = b
		b += blockSize
	}
	InsertSort(array[a:n])     // 防止元素的个数不是 blockSize 的倍数

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			merge(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		// 防止元素的个数不是 2*blockSize 的倍数
		if m := a + blockSize; m < n {
			merge(array, a, m, n)
		}
		blockSize *= 2
	}
}

// 原地归并操作
func merge(array []int, begin, mid, end int) {
	// 三个下标，将数组 array[begin, mid] 和 array[mid, end-1]进行原地归并
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从 j 向右移动，找到第一个 array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		// 进行手摇翻转，将 array[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		rotation(array, i, j-step, j-1)
		i = i + step
	}
}

// 手摇算法
func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid) // 前面逆序
	reverse(array, mid, r) // 后面逆序
	reverse(array, l, r)   // 整体逆序
}

// 逆序
func reverse(array []int, l, r int) {
	for l < r {
		// 左右互换交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}

func main() {
	list := []int{5}
	MergeSort(list, len(list))
	fmt.Println(list)

	list1 := []int{5, 9}
	MergeSort(list1, len(list1))
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MergeSort(list2, len(list2))
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 45, 67, 2, 5, 24, 56, 34, 24, 56, 2, 2, 21, 4, 1, 4, 7, 9}
	MergeSort(list3, len(list3))
	fmt.Println(list3)
}
