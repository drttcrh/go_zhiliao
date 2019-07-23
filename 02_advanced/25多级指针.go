package main

import "fmt"

func main() {

	a := 10
	// 一级指针，指向变量的地址
	p := &a
	fmt.Println(p)
	fmt.Printf("%T\n", p)		// *int

	// 二级指针，指向一级指针的地址
	p2 := &p
	fmt.Printf("%T\n", p2)		// **int
}
