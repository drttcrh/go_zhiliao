package _1_basic

import "fmt"

func main() {
	var ch byte
	var str string
	// 字符
	// 1.字符只能用单引号
	// 2.字符，往往都只是一个字符，除了 \n, \t 转义字符
	ch = 'a'
	fmt.Println(ch)

	// 字符串
	// 1. 字符串必须用双引号
	// 2. 字符串可以有一个或多个字符组成
	// 3. 字符串都是隐藏了一个结束符 '\0'
	str = "a"
	fmt.Println(str)

	// 获取字符串中的字符
	str1 := "hello world"
	fmt.Println(str1[0], str1[1])	// 104 101 获取字符时获取的是字符对应的 ascii 码
	fmt.Printf("%c, %c\n", str1[0], str1[1])	// h, e
}
