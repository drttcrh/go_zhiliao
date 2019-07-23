package main

import "fmt"

// 定义接口 一般接口名以 er 结尾，根据接口实现功能
type Humener12 interface {
	// 方法
	sayHi()
}

type student12 struct {
	name string
	age int
	score int
}

func (s *student12) sayHi() {
	fmt.Printf("大家好，我是%s，我今年%d，我的分数是%d\n", s.name, s.age, s.score)
}

type teacher12 struct {
	name string
	age int
	subject string
}

func (t *teacher12) sayHi() {
	fmt.Printf("大家好，我是%s，我今年%d，我的科目是%s\n", t.name, t.age, t.subject)
}

// 多态的实现，将接口作为函数参数
func sayHello(h Humener12)  {
	h.sayHi()
}

func main() {

	// 多态的实现
	stu := student12{"loedan", 13, 89}
	// 调用多台函数
	sayHello(&stu)

	tea := teacher12{"wang", 23, "golang"}
	sayHello(&tea)
}
