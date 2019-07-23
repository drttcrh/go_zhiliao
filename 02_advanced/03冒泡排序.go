package main

import "fmt"

func main() {
	arr := [10] int {4, 6, 3, 1, 7, 5, 9, 0, 2, 8}

	// 冒泡排序
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			}
		}
	}
	fmt.Println(arr)
}
