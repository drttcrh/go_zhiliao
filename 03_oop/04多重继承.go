package main

import "fmt"

type person4 struct{
	name string
	age int
	sex string
}

// 父类
type Person4 struct {
	id int
	addr string
}

// 子类
type Student4 struct {

	// 多重继承
	person4		// 匿名字段
	Person4		// 匿名字段
	score int
}

func main() {
	// 多继承实例并继承
	var s Student4
	s.person4.name = "loedan"
	s.person4.age = 23
	s.person4.sex = "female"
	s.Person4.id = 1
	s.Person4.addr = "jiangxi"
	s.score = 199
	fmt.Println(s)		// {{loedan 23 female} {1 jiangxi} 199}

	// 自动推导初始化
	s1 := Student4{
		person4{"aha", 13, "male"},
		Person4{13, "shanghai"},
		23,
	}
	fmt.Println(s1)
}
