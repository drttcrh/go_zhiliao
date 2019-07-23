package _1_basic

import "fmt"

func main() {
	// 1. iota是常量自动生成器，每个一行，自动加1
	// 2. iota给常量赋值使用
	const (
		a = iota	// 0
		b = iota	// 1
		c = iota	// 2
	)
	fmt.Println(a, b, c)

	// 3.iota遇到 const， 重置为 0
	const d = iota	// 0
	const e = iota	// 0
	fmt.Println(d, e)

	// 4.可以只写一个 iota
	const (
		f = iota
		g
		h
	)
	fmt.Println(f, g, h)

	// 5.如果常量写在同一行，值是相等的
	const (
		i = iota	// 0
		j, k, l = iota, iota, iota	// 1, 1, 1
		m = iota	// 2
	)
}
