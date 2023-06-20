/*
第一轮迭代：从第一个数开始，依次比较相邻的两个数，如果前面一个数比后面一个数大，那么交换位置，直到处理到最后一个数，最后的这个数是最大的。
第二轮迭代：因为最后一个数已经是最大了，现在重复第一轮迭代的操作，但是只处理到倒数第二个数。
......

第一次比较的次数为： N-1 次
第二次比较的次数为： N-2 次，因为排除了最后的元素
第三次比较的次数为： N-3 次，因为排除了后两个元素
......
比较次数：1 + 2 + 3 + ... + (N-1) = (N^2 - N)/2

冒泡排序可以算是最慢的排序算法了 -.-
*/
package main

import "fmt"

// i = n-1, j最大为 n-2，可以访问到的切片下标最大为 n-1 满足
// i = 1,j最大为 0，可以访问到的切片下标最大为 1 满足
func BubbleSort(list []int) {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		// 标记在一轮中有没有交换过
		flag := false
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				// temp := list[j]
				// list[j] = list[j+1]
				// list[j+1] = temp
				list[j], list[j+1] = list[j+1], list[j] // 在go语言中允许这样做，在python中这样也是允许的
				flag = true
			}
		}
		// 在这一轮中如果没有交换，说明排序完了，直接返回
		if !flag {
			return
		}
	}
}

func main() {
	list1 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(list1)
	fmt.Println(list1)

	list2 := []int{1, 1, 1, 6, 8, 14, 6, 49, 25, 4, 6, 1}
	BubbleSort(list2)
	fmt.Println(list2)
}
