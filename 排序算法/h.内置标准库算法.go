package main

import (
	"fmt"
	"sort"
)

func InnerSort() {
	list := []struct {
		Name string
		Age  int
	}{
		{"A", 75},
		{"B", 4},
		{"C", 5},
		{"D", 5},
		{"E", 2},
		{"F", 5},
		{"G", 5},
	}
	// SliceStable 是稳定排序，使用了插入排序和归并排序  SliceStable(x any, less func(i, j int) bool)
	sort.SliceStable(list, func(i, j int) bool { return list[i].Age < list[j].Age })
	fmt.Println(list)

	list2 := []struct {
		Name string
		Age  int
	}{
		{"A", 75},
		{"B", 4},
		{"C", 5},
		{"D", 5},
		{"E", 2},
		{"F", 5},
		{"G", 5},
	}
	// 是一般的排序，不追求稳定排序，使用了快速排序  Slice(x any, less func(i, j int) bool)
	sort.Slice(list2, func(i, j int) bool { return list2[i].Age < list2[j].Age })
	fmt.Println(list2)
	/*  采用了四种优化
	第一种是递归时小数组转为插入排序
	第二种是当递归超过 2*ceil(log(n+1)) 层，内部会转为堆排序
	第三种是使用了中位数基准数
	第四种使用了三向切分
	*/
}

func main() {
	InnerSort()
}
