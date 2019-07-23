package main

import "fmt"

func mapParmas(m map[int]string) {
	m[4] = "四哥"
	fmt.Println(m)
}

func main() {
	m := map[int]string{1: "西施", 2: "貂蝉", 3: "大乔"}
	fmt.Println(m)
	// 引用传递，函数中对行参进行修改，会影响实参的内容
	mapParmas(m)
	fmt.Printf("%p\n", m)
	fmt.Println(m)
}
