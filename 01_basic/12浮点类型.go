package main

import "fmt"

func main() {
	// 浮点型数据分为单精度（float32）和 双精度（float64）
	var a float64 = 123.3242323323223
	fmt.Println(a)	// 123.3242323323223

	var b float32 = 123.3242323323223
	fmt.Println(b)	// 123.324234

	// 自动推导创建的浮点型变量，默认为 float64
	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T\n", c)		// float64
}
