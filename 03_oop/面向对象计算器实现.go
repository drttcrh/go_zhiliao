package main

import "fmt"

type Opter interface {
	// 方法声明
	Result() int
}

type Operate struct {
	num1 int
	num2 int
}

type Add struct {
	Operate
}

func (a *Add) Result() int {
	return a.num1 + a.num2
}

type Sub struct {
	Operate
}

func (s *Sub) Result() int {
	return s.num1 - s.num2
}

func Result(o Opter) {
	v := o.Result()
	fmt.Println(v)
}

func main() {
	// 面向对象实现计算器
	var o Opter
	var a Add = Add{Operate{10, 20}}
	o = &a
	result := o.Result()
	fmt.Println(result)

	// 多态的方式实现
	var s Sub = Sub{Operate{20, 10}}
	Result(&s)
}
