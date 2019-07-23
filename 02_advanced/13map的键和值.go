package main

import "fmt"

func main() {
	m := make(map[string][3]int)	// 创建一个key为string类型，值为数组类型的map
	m["loedan"] = [3]int{100, 23, 43}
	m["harden"] = [3]int{23, 43, 54}
	m["kobe"] = [3]int{1, 2, 3}
	fmt.Println(m)

	// 取值
	for k, v := range m {
		fmt.Println(k, v)
	}


	m1 := make(map[int]string)
	m1[1] = "loedan"
	m1[2] = "harden"
	m1[3] = "macgrady"
	fmt.Println(len(m1))	// 3
	fmt.Println(m1[4])		// 在 map 中如果没有不存在要找的key，会返回map值类型的默认值
	fmt.Println(len(m1))	// 3

	// 判断 map 中某一个 key 是否存在
	value, ok := m1[1]
	fmt.Println(ok, value)

	value1, ok := m1[5]
	fmt.Println(ok, value1)

	// 删除 map 中的元素
	delete(m1, 3)
	fmt.Println(m1)
}
