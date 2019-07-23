package main

import "fmt"

// 数组作为函数参数
// 冒泡排序
func BubbleSort(arr [10] int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			}
		}
	}
	fmt.Println(arr)	// [-2 0 1 2 4 5 6 7 8 9]
}

func main() {
	arr := [10] int {2, 4, 5, 9, 0, -2, 6, 8, 1, 7}

	// 数组作为函数参数是值传递
	// 行参和实参在内存中是不同的地址
	BubbleSort(arr)

	fmt.Println(arr)	// [2 4 5 9 0 -2 6 8 1 7]，数组内容没有的到改变
}
