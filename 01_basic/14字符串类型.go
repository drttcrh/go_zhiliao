package main

import "fmt"

func main() {
	var str1 string
	str1 = "ab"
	fmt.Println(str1)

	// 自动推导类型
	str2 := "loedan"
	fmt.Println(str2)
	fmt.Printf("%T\n", str2)

	ch := 'a'
	str3 := "a"	// 'a''\0'	\0 表示字符串结束标志
	fmt.Println(ch)
	fmt.Println(str3)
	fmt.Printf("%T\n", ch)
	fmt.Printf("%T\n", str3)

	// len函数，计算字符串中字符的个数
	// 在 go 语言中，一个汉子占三个字符
	fmt.Println(len(str3))

	// 字符串拼接
	str4 := "hello"
	str5 := "world"
	str6 := str4 + " " + str5
	fmt.Println(str6)
}
