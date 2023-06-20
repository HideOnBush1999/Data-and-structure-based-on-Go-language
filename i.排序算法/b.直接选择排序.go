/*
第一轮迭代，从第一个数开始，左边到右边进行扫描，找到最小的数，与数列里的第一个数交换位置。
第二轮迭代，从第二个数开始，左边到右边进行扫描，找到最小的数，与数列里的第二个数交换位置。
......
*/

package main

import "fmt"

func SelectSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		min := list[i] // 最小值先默认赋值为当前切片值
		index := i     // 记录最小值的索引
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				index = j
			}
		}
		if index != i { // 如果当前切片值不是最小值
			list[i], list[index] = list[index], list[i]
		}
	}
}

// 我们每一轮，除了找最小数之外，还找最大数
// 然后分别和前面和后面的元素交换，这样循环次数减少一半
func SelectGoodSort(list []int) {
	n := len(list)

	// 只需要循环一半
	for i := 0; i < n/2; i++ {
		// 每轮找出最大值和最小值
		maxIndex := i
		minIndex := i
		for j := i + 1; j < n-i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		/*
			// 下面思想上简单粗暴，但是会出现逻辑错误
			// 比如说当你最小值的位置在 n-i-1，那么第一个判断之后，最小值的位置发生了变化，但是minIndex不变
			if maxIndex != n-i-1 {
				list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			}
			if minIndex != i {
				list[i], list[minIndex] = list[minIndex], list[i]
			} */

		// 所以需要同时考虑最大索引和最小索引
		if maxIndex == i && minIndex == n-i-1 { // 最大值在开头，最小值在结尾
			list[maxIndex], list[minIndex] = list[minIndex], list[maxIndex]
		} else if maxIndex == i && minIndex != n-i-1 { // 最大值在开头，最小值不在结尾
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1] // 要先交换最大值
			list[i], list[minIndex] = list[minIndex], list[i]
		} else { // 最大值不在开头
			list[i], list[minIndex] = list[minIndex], list[i] // 要先交换最小值
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	SelectGoodSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	SelectGoodSort(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectGoodSort(list3)
	fmt.Println(list3)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
	SelectGoodSort(list4)
	fmt.Println(list4)
}
