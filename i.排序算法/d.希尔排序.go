/*
先取一个小于 N 的整数 d1，将位置是 d1 整数倍的数们分成一组，对这些数进行直接插入排序。
接着取一个小于 d1 的整数 d2，将位置是 d2 整数倍的数们分成一组，对这些数进行直接插入排序。
接着取一个小于 d2 的整数 d3，将位置是 d3 整数倍的数们分成一组，对这些数进行直接插入排序。
...
直到取到的整数 d=1，接着使用直接插入排序。
*/
package main

import "fmt"

// 增量序列折半的希尔排序
func ShellSort(list []int) {
	// 数组长度
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，对位置是 step 倍数的元素进行直接插入排序
		for i := step; i < n; i += step {
			deal := list[i] // 待排序的数
			j := i - step   // 待排序的数的左边最近一个数的下标
			for ; j >= 0 && deal < list[j]; j -= step {
				list[j+step] = list[j]
			}
			list[j+step] = deal
			printList(list)
		}
	}
}

func printList(list []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", list[i])
	}
	fmt.Printf("\n")
}

func main() {
	list := []int{5}
	ShellSort(list)
	fmt.Println(list)

	list1 := []int{7, 3, 5, 9, 2, 0, 8, 6}
	ShellSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	ShellSort(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
	ShellSort(list3)
	fmt.Println(list3)
}
