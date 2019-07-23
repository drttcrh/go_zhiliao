package main

import "fmt"

// 数组指针作为函数参数
func arrayPointerParams(p *[5]int)  {
	(*p)[0] = 32
}

func main() {
	arr := [3]int{1, 2, 3}
	// 定义数组指针
	var p *[3]int
	// 指针和数组建立关系
	p = &arr
	fmt.Println(*p)


	// 通过指针操作数组
	// (*数组指针变量)[下标] = 值
	(*p)[0] = 22
	fmt.Println(*p)
	// 直接赋值
	p[1] = 33
	fmt.Println(*p)


	arr1 := [5]int{1, 2, 3, 4, 5}
	// 指针变量和要存储数据类型要相同
	p1 := &arr1
	p2 := &arr1[0]
	// p1 和 p2 在内存中指向相同的地址
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)
	// p1 和 p2 在内存中的指针类型不同
	fmt.Printf("%T\n", p1)	// *[5]int	数组指针
	fmt.Printf("%T\n", p2)	// *int		变量指针


	// 数组指针作为函数参数
	arr2 := [5]int{1, 2, 3, 4, 5}
	arrayPointerParams(&arr2)
	fmt.Println(arr2)
}
