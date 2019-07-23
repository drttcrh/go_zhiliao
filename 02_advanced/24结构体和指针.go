package main

import "fmt"

type person struct {
	id int
	name string
	age int
}

// 结构体指针作为函数参数
func structSliceParams(p *person) {
	p.name = "harden"
}

func main() {
	// 创建一个结构体实例
	var per person = person{1, "loedan", 13}
	fmt.Println(per)
	fmt.Printf("%p\n", &per)		// 0xc00000c080


	// 定义指针接收结构体变量地址
	p := &per
	fmt.Printf("%T\n", p)		// *main.person 结构体指针类型

	// 通过指针间接修改结构体成员的值
	(*p).age = 23
	fmt.Println(per)		// {1 loedan 23}
	// 指针直接操作结构体成员
	p.id = 12
	fmt.Println(per)		// {12 loedan 23}


	// 结构体指针作为函数参数
	structSliceParams(p)
	fmt.Println(per)		// {12 harden 23}



	// 结构体数组指针
	arr := [3]person{
		{1, "loedan", 23},
		{2, "harden", 11},
		{3, "macgrady", 121},
	}
	fmt.Println(arr)		// [{1 loedan 23} {2 harden 11} {3 macgrady 121}]
	// 指向结构体数组的指针
	p1 := &arr
	// 通过指针修改结构体数组的元素内容
	p1[0].age = 18
	fmt.Println(arr[0])		// {1 loedan 18}



	// map 类型数组结构体指针
	m := make(map[int]*[3]person)
	fmt.Printf("%T\n", m)		// map[int]*[3]main.person
	// 赋值
	m[1] = new([3]person)
	m[1] = &[3]person{
		{1, "loedan", 23},
		{2, "harden", 11},
		{3, "macgrady", 121},
	}
	m[2] = new([3]person)
	m[3] = &[3]person{
		{1, "xixixi", 23},
		{2, "hahaha", 11},
		{3, "wowowo", 121},
	}

}
