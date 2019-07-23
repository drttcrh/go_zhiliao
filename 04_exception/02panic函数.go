package main

import "fmt"

// 数组越界异常
func crossing(i int) {
	var arr [3]int
	arr[i] = 999
	fmt.Println(arr)
}

func main() {
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	// panic("hello")
	fmt.Println("hello")

	crossing(4)	// panic: runtime error: index out of range
}
