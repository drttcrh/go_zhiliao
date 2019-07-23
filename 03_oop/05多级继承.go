package main

import "fmt"

type humen struct {
	id int
	name string
}

// 继承 humen
type person5 struct {
	humen
	age int
	sex string
}

// 继承 person5
type student5 struct {
	person5
	score int
}

func main() {
	// 实例化多级继承成员
	s := student5{
		person5{
			humen{1, "loedan"},
			18,
			"male",
		},
		122,
	}
	fmt.Println(s)		// {{{1 loedan} 18 male} 122}

	// 初始化并赋值
	var s1 student5
	s1.id = 1
	s1.name = "harden"
	s1.age = 23
	s1.sex = "female"
	s1.score = 199
	fmt.Println(s1)		// {{{1 harden} 23 female} 199}
}
