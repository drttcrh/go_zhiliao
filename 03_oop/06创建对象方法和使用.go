package main

import "fmt"

type cat struct {
	name string
	age int
}

type dog struct {
	name string
	age int
}

// 方法定义
//func (对象) 方法名(参数列表) (返回值列表) {
//	coding
//}

// 结构体作为接收者
func (c cat) show() {
	fmt.Println("this is cat's show func")
	fmt.Printf("我是%s\n", c.name)
}

func (d dog) show() {
	fmt.Println("this is dog's show func")
}

func show() {
	fmt.Println("this is show func")
}

func main() {
	var c cat
	c.name = "cat"
	c.age = 32
	c.show()

	var d dog
	d.name = "dog"
	d.age = 322
	d.show()

	// 对象的方法名 和 包的函数名 可以重复
	show()
}
