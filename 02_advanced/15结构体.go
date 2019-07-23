package main

import "fmt"

// 结构体定义在函数外部
// 定义: type 结构体名 struct {
// 		结构体成员列表
// 		成员名 数据类型
// }
// 学生结构
type stu struct {
	id int
	name string
	sex string
	age int
	addr string
}

// 实例化结构体
func instantiationStruct() {

	// 方式一
	// 顺序初始化，每个成员都必须初始化
	var s stu = stu{
		1,
		"loedan",
		"female",
		18,
		"i don't know",
	}
	fmt.Println(s)

	// 方式二
	// 自动推导初始化
	s1 := stu{
		name: "sige",
		addr: "大婚头",
		id: 2,
		sex: "male",
		age: 13,
	}
	fmt.Println(s1)

	// 方式三
	var s2 stu
	fmt.Println(s2)		// {0   0 }
	// 给元素赋值
	s2.name = "jordan"
	s2.age = 99
	s2.id = 3
	s2.addr = "hahaha"
	s2.sex = "male"
	fmt.Println(s2)
}

// 结构体的内存地址
func structRam() {

	var s stu = stu{
		1,
		"loedan",
		"female",
		18,
		"i don't know",
	}
	// 结构体名本身指向第一个成员的内存地址
	fmt.Printf("%p\n", &s)		// 0xc00007e140
	fmt.Printf("%p\n", &s.id)		// 0xc00007e140
}

// 结构体比较
func structContrast() {
	var s stu = stu{
		1,
		"loedan",
		"female",
		18,
		"i don't know",
	}
	s1 := stu{
		2,
		"sige",
		"male",
		13,
		"大婚头",
	}
	s2 := stu{
		2,
		"sige",
		"male",
		13,
		"大婚头",
	}
	fmt.Println(s == s1)	// false
	fmt.Println(s1 == s2)	// true
}

// 同类型的两个结构体可以相互赋值
func structValueSet() {
	var s stu = stu{
		1,
		"loedan",
		"female",
		18,
		"i don't know",
	}
	var s1 stu
	s1 = s
	fmt.Println(s1)		//
}

func main() {
	// 实例化结构体
	instantiationStruct()

	// 结构体的内存地址
	structRam()

	// 结构体比较
	structContrast()

	// 结构体与结构体之间的相互赋值
	structValueSet()
}
