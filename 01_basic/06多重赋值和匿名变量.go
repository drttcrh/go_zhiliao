package _1_basic

import "fmt"

func test() (a, b, c int){
	return 1, 2, 3
}

func main() {
	//a := 10
	//b := 20
	//c := 30
	// 多重赋值
	a, b := 10, 20
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)	// 20, 10

	// _ 匿名变量，丢弃数据不处理
	tmp, _ := 3, 4
	fmt.Println(tmp)

	_, c, d := test()
	fmt.Println(c, d)
}
