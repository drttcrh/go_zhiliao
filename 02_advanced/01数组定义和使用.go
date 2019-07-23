package main

import "fmt"

func main() {
	// 数组是指一系列同一类型数据的集合
	// 数组定义
	// var 数组名 [长度] 类型
	var arr [10] int
	fmt.Println(arr)

	// 数组赋值
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	// arr[-1] = -1	// 报错
	fmt.Println(arr)
	// 通过循环快速给数组赋值
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	// 通过循环快速对数组取值
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for _, v := range arr{
		fmt.Println(v)
	}

	// 注意
	// var arr [10] int	// int 类型数组元素默认值为 0
	// var arr [10] float63 // float 类型数组元素默认值为 0
	// var arr [10] bool	// bool 类型数组元素默认值为 false


	// 数组初始化
	// 1、全部初始化
	var arr1 [5] int = [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 自动推导
	arr2 := [5] int {1, 2, 3, 4}
	fmt.Println(arr2)	// [1 2 3 4 0]	没有初始化的元素，默认为 0

	// 指定元素的初始化
	arr3 := [5] int {2 : 10, 4 : 90}	// 2 : 10, 2 表示数组下标，10 表示值
	fmt.Println(arr3)	// [0 0 10 0 90]

	// 通过 ... 初始化数组的长度定义
	arr4 := [...] int {1, 2, 3}
	fmt.Println(arr4)	// [1 2 3]
	fmt.Println(len(arr4))	// [1 2 3]


	// 常见问题
	// 数组长度必须是常量
	// 数组下标不能越界
	// go 语言数组没有 -1 等类型下标，只有正整数
	arr5 := [5] int {1, 2, 3, 4, 5}
	// %p 打印元素在内存中的地址
	// 数组名表示整个数组，数组名地址就是数组第一个元素的地址
	fmt.Printf("%p\n", &arr5)		// 0xc00001c0c0
	fmt.Printf("%p\n", &arr5[0])	// 0xc00001c0c0
	fmt.Printf("%p\n", &arr5[1])	// 0xc00001c0c8
	fmt.Printf("%p\n", &arr5[2])	// 0xc00001c0d0
	fmt.Printf("%p\n", &arr5[3])	// 0xc00001c0d8
	fmt.Printf("%p\n", &arr5[4])	// 0xc00001c0e0
}
