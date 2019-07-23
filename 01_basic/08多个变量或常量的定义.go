package _1_basic

import "fmt"

func main() {
	// 不同类型变量的定义
	var a int = 1
	var b float64 = 2.0
	fmt.Println(a, b)

	// 自动推导类型
	var (
		c = 1
		d = 2.0
	)
	fmt.Println(c, d)

	// 常量的定义
	const e int = 1
	const f float64 = 3.04
	fmt.Println(e, f)

	// 自动推导
	const (
		g int = 3
		h float64 = 3.14
	)
	fmt.Println(g, h)

	const (
		i = 3
		j = 3.13
	)
	fmt.Println(i, j)
	fmt.Printf("%T, %T", i, j)
}
