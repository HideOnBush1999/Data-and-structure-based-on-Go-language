package main

import "fmt"

func f1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f1(n-1) + f1(n-2)
}

// 尾递归
func f2(n int, a1, a2 int) int {
	if n == 1 {
		return a1
	}
	return f2(n-1, a2, a1+a2)
	// a1 变成了下一个数 a2
	// a2 变成了下一个数 a1 + a2
}

func main() {
	fmt.Println(f1(6))
	fmt.Println(f2(6, 1, 1)) // 斐波那契数列第一项和第二项都为 1
}
