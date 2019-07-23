package main

import "fmt"

func main() {

	var i int
	i = 100
	fmt.Println(i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%p\n", &i)	// 获取变量在内存中的地址 0xc00001a058

	// 定义指针变量
	var p *int
	p = &i
	fmt.Println(p)

	// 通过指针修改变量的值
	*p = 80
	fmt.Println(*p)

	// 注意
	// 1、指针类型默认值为 nil
	// 空指针
	var p1 *int
	fmt.Println(p1)		// <nil>

	// 野指针，指针指向一个未知的空间，会报错，程序不允许出现野指针
	var p2 *int
	//*p2 = 32
	fmt.Println(p2)


	// new 函数
	var p3 *int
	p3 = new(int)	// new() 在内存中开辟一块对应类型的空间
	*p3 = 22
	fmt.Println(*p3)

	// 自动推导类型
	p4 := new(int)
	*p4 = 999
	fmt.Println(*p4)

}
