package main

import "fmt"

type person9 struct {
	id int
	name string
	age int
}

type student9 struct {
	person9
	class int
}

func (p *person9) PrintInfo() {
	fmt.Printf("编号%d\n", p.id)
	fmt.Printf("姓名%s\n", p.name)
	fmt.Printf("年纪%d\n", p.age)
}

// 重写父类的方法
func (s *student9) PrintInfo()  {
	fmt.Println("student:", *s)		// student: {{2 aha 23} 3}
}

func main() {
	p := person9{1, "loedan", 123}
	p.PrintInfo()

	s := student9{person9{2, "aha", 23}, 3}
	s.PrintInfo()			// 就近原则
	s.person9.PrintInfo()	// 调用父类中的同名方法
}
