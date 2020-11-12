package main

import "fmt"

func main() {

	a := 10
	b := "abs"
	c := 'a'
	d := 2.3452

	// %T 打印变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)	// int, string, int32, float64

	// %d 整形格式输出
	// %s 字符串格式输出
	// %c 字符格式输出
	// %f 浮点型格式输出, %.2f 表示只取小数点后 2 位，最后会四舍五入
	fmt.Printf("%d, %s, %c, %f, %.2f\n", a, b, c, d, d)	// 10, abs, a, 2.345200, 2.35

	// %v 自动格式匹配输出
	// %v 输出字符类型时输出的是字符的 ascii 码
	fmt.Printf("%v, %v, %v, %v\n", a, b, c, d)	// 10, abs, 97, 2.3452
}
