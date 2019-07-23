package main

import "fmt"

func main() {
	// 一维数组定义 One-Dimensional Array
	var arr [5] int
	fmt.Println(arr)	// [0 0 0 0 0]

	// 二维数组定义 double-Dimensional Array
	var arr1 [2][3] int	// 表示一个两行三列的二维数组
	fmt.Println(arr1)	// [[0 0 0] [0 0 0]]

	// 二维数组赋值
	arr1[0][1] = 234	// 给第一行第二个元素赋值
	arr1[1][2] = 23490	// 给第二行第三个元素赋值
	fmt.Println(arr1)	// [[0 234 0] [0 0 23490]]

	// 取出二维数组中所有的值
	for _, v := range arr1 {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}


	// 二维数组定义并初始化
	// 1、 全部初始化
	arr2 := [3][4] int { {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12} }
	fmt.Println(arr2)
	// 2、部分初始化
	arr3 := [3][4] int { {1, 2, 3, 4}, {5, 6, 7, 8} }
	fmt.Println(arr3)
}
