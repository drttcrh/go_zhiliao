package main

import "fmt"

func main() {
	// 接口类型可以接收任意类型的数据


	// 定义一个空接口
	var i interface{}
	fmt.Println(i)					// <nil>
	fmt.Printf("%T\n", i)		// <nil>

	i = 10
	fmt.Println(i)					// 10
	fmt.Printf("%T\n", i)		// int

	i = 3.14
	fmt.Println(i)					// 3.14
	fmt.Printf("%T\n", i)		// float64

	i = true
	fmt.Println(i)					// true
	fmt.Printf("%T\n", i)		// bool

	i = "golang"
	fmt.Println(i)					// golang
	fmt.Printf("%T\n", i)		// string


	// 空接口类型的切片
	var s []interface{}
	fmt.Printf("%T\n", s)		// []interface {}
	s = append(s, 10, 2.23, true, "golang", [2]int{1, 2}, []string{"loedan", "aha"}, test14)
	fmt.Println(s)
}

func test14()  {
	fmt.Println("test14")
}
