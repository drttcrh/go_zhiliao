package main

import "fmt"

func main() {
	// type 给数据类型取别名

	// 给 int64 起一个别名 bigint
	type bigint int64
	var a bigint = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)		// main.bigint

	// 取一组类型别名
	type (
		long int64
		char byte
	)
	var l long = 32
	var c char = 'a'
	fmt.Println(l, c)
}
