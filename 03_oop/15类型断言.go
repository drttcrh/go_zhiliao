package main

import "fmt"

// 断言用函数
func comma()  {
	fmt.Println("abcdefg")
}
func main() {
	var i interface{}
	i = 10

	// 断言
	// value, ok := element.(T)
	// 值，值的判断 := 接口变量.(数据类型)

	_, ok := i.(int)
	if ok {
		fmt.Println("int")
	} else {
		fmt.Println("unknown")
	}


	// 切片类型接口断言
	var j []interface{}
	j = append(j, 1, 3.14, "aaa", comma)
	for _, v := range j {
		switch value := v.(type) {
		case int:
			fmt.Println("整形数据:", value)
		case string:
			fmt.Println("字符串", value)
		case func():
			value()
		}
	}
}
