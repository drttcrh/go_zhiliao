package main

import "fmt"

// 定义接口 一般接口名以 er 结尾，根据接口实现功能
type Humener interface {
	// 方法
	sayHi()
}

type student11 struct {
	name string
	age int
	score int
}

func (s *student11) sayHi() {
	fmt.Printf("大家好，我是%s，我今年%d，我的分数是%d\n", s.name, s.age, s.score)
}

type teacher11 struct {
	name string
	age int
	subject string
}

func (t *teacher11) sayHi() {
	fmt.Printf("大家好，我是%s，我今年%d，我的科目是%s\n", t.name, t.age, t.subject)
}

func main() {
	// 接口是一种数据类型，可以接受满足对象的信息
	// 接口是虚的，方法是实的
	// 接口定义方法，结构体实现方法
	// 接口定义的方法，结构体中必须全部实现
	var h Humener
	stu := student11{"小明", 23, 43}
	//stu.sayHi()
	h = &stu
	h.sayHi()

	tea := teacher11{"teacher wang", 23, "golang"}
	//tea.sayHi()
	h = &tea
	h.sayHi()
}
