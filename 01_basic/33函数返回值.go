package main

import "fmt"

// 单个返回值
func oneReturnValFunc(a, b int) int {
	return a + b
}

// 多个返回值
func multiReturnValFunc(a string, b int) (int, string) {
	return b, a
}

func main() {
	// 单个返回值
	result := oneReturnValFunc(1, 2)
	fmt.Println(result)

	// 多个返回值
	a, b := multiReturnValFunc("loedan", 2)
	fmt.Println(a, b)

	_, d := multiReturnValFunc("loedan", 4)	// 丢弃函数返回值
	fmt.Println(d)
}
