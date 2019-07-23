package main

import "fmt"

// 结构体嵌套结构体
// 父类
type Person2 struct {
	id int
	name string
	age int
}

// 子类
type Student2 struct {
	Person2	// 匿名字段
	name string
	score int
}

func main() {
	var s Student2

	// 子类和父类结构体有相同的成员名，默认赋值给子类，采用就近原则
	s.name = "loedan"
	s.Person2.name = "aha"
	fmt.Println(s)		// {{0 aha 0} loedan 0}
}
