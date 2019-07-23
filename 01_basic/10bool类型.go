package _1_basic

import "fmt"

func main() {
	// 1.声明变量 没有初始化，默认为 false
	// 2.不二类型不接受其他类型的赋值，不支持自动或强制的类型转换
	var a bool
	fmt.Println(a)

	a = false
	fmt.Println(a)

	// 3.自动推导类型
	var b = true
	fmt.Println(b)

	c := false
	fmt.Println(c)

	d := (1 == 2)
	fmt.Println(d)

}
