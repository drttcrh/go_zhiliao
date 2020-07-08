package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// uint 无符号整形 int 有符号整形
	var a int = -10
	fmt.Println(a)

	b := 20
	fmt.Println(b)
	// %T 查看变量类型
	fmt.Printf("%T\n", b)
	// unsafe.Sizeof()	查看字段长度
	fmt.Println(unsafe.Sizeof(b))

	var c int32 = 10
	fmt.Println(c)
	// %T 查看变量类型
	fmt.Printf("%T\n", c)
	// unsafe.Sizeof()	查看字段长度
	fmt.Println(unsafe.Sizeof(c))
}
