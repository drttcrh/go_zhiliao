package main

import "fmt"

type person8 struct {
	id int
	name string
	age int
}

type student8 struct {
	person8
	class int
}

func (p *person8) PrintInfo() {
	fmt.Printf("编号%d\n", p.id)
	fmt.Printf("姓名%s\n", p.name)
	fmt.Printf("年纪%d\n", p.age)
}

func main() {
	p := person8{1, "loedan", 123}
	p.PrintInfo()

	s := student8{person8{2, "aha", 23}, 3}
	// 调用父类的方法的两种方式
	s.PrintInfo()
	s.person8.PrintInfo()
}
