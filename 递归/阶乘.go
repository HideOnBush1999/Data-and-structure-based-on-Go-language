package main

import "fmt"

func Rescuvie(n int) int {
	if n == 1 {
		return 1
	}
	return n * Rescuvie(n-1)
}

// 尾递归
func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

func main() {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))
}
