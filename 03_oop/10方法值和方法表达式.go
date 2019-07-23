package main

import "fmt"

type person10 struct {
	id int
	name string
	age int
}

func (p person10) PrintInfo() {
	fmt.Printf("%p, %v\n", &p, p)
}

func (p *person10) PrintInfo1() {
	fmt.Printf("%p, %v\n", p, *p)
}

func main() {
	p := person10{1, "loedan", 123}
	p.PrintInfo()		// 0xc00000c080, {1 loedan 123}
	p.PrintInfo1()		// 0xc00000c060, {1 loedan 123}

	fmt.Println(p.PrintInfo)		// 0x10951d0
	fmt.Println(p.PrintInfo1)		// 0x1095340

	fmt.Printf("%T\n", p.PrintInfo)		// func()
	fmt.Printf("%T\n", p.PrintInfo1)		// func()


	// 对象相同，但是函数内容不同，不能赋值
	// 方法值，隐士传递，隐藏的是接收者
	var pfunc1 func()
	pfunc1 = p.PrintInfo
	pfunc1()
	fmt.Printf("%T\n", pfunc1)


	// 方法表达式
	pfunc2 := person10.PrintInfo
	pfunc2(p)
}
