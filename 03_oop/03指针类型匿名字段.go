package main

import "fmt"

// 结构体嵌套结构体
// 父类
type Person3 struct {
	id int
	name string
	age int
}

// 子类
type Student3 struct {
	*Person3	// 指针类型匿名字段
	score int
}

func main() {
	var s Student3
	s.score = 100
	fmt.Println(s)		// {<nil>  100}
	s.Person3 = new(Person3)
	s.Person3.id = 23
	s.Person3.name = "loedan"
	s.Person3.age = 12
	fmt.Println(s)		// {0xc0000a0000 100}
	fmt.Println(s.id, s.name, s.age)		// 23 loedan 12
}
