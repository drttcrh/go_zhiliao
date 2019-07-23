package main

import "fmt"

func main() {
	// 切片的内置函数
	// append()
	// copy()	拷贝

	// 拷贝
	s := [] int {1, 2, 3, 4, 5}
	s1 := make([] int, 5)
	fmt.Println(s1)	// [0 0 0 0 0]
	copy(s1, s)		// 将 s 中的数据拷贝到 s1 中
	fmt.Println(s1)	// [1 2 3 4 5]

	// 使用 copy 进行拷贝，在内存中存储两个独立的切片内容
	fmt.Printf("%p\n", s)		// 0xc00001c060
	fmt.Printf("%p\n", s1)	// 0xc00001c090
}
