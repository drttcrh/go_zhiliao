package main

import "fmt"

// 定义接口 一般接口名以 er 结尾，根据接口实现功能
type Humener13 interface {	// 子集
	// 方法
	sayHi()
}

type Personer13 interface {	// 超集
	Humener13	// 继承
	sing(string)
}

type student13 struct {
	name string
	age int
	score int
}

func (s *student13) sayHi() {
	fmt.Printf("大家好，我是%s，我今年%d，我的分数是%d\n", s.name, s.age, s.score)
}

func (s *student13) sing(str string) {
	fmt.Println(str)
}

func main() {
	// 定义接口类型变量
	var h Humener13
	stu := student13{"loedan", 21, 89}
	h = &stu
	h.sayHi()

	// 定义接口类型变量
	var p Personer13
	p = &stu
	p.sing("ahahahaha")


	// 接口转换
	var h1 Humener13
	var p1 Personer13
	var stu1 student13 = student13{"loedan", 23, 43}

	// 超集可赋值给子集，子集无法赋值给超集
	p1 = &stu1
	h1 = p1
	h1.sayHi()
}
