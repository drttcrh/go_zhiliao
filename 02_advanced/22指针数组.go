package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30

	// 指针数组，数组元素为指针类型
	var arr [3]*int = [3]*int{&a, &b, &c}
	fmt.Println(arr)	// [0xc00001a060 0xc00001a068 0xc00001a070]

	*arr[1] = 200
	fmt.Println(b)	// 200

	// 循环读取
	for i := 0; i < len(arr); i++ {
		fmt.Println(*arr[i])
	}


	// 数组中存储数组的指针
	d := [3]int{1, 2, 3}
	e := [3]int{4, 5, 6}
	f := [3]int{7, 8, 9}
	var arr1 [3]*[3]int = [3]*[3]int{&d, &e, &f}
	fmt.Println(arr1)	// [0xc000084040 0xc000084060 0xc000084080]

	// 修改指针的内容
	(*arr1[2])[2] = 999
	fmt.Println(f)	// [7 8 999]
}
