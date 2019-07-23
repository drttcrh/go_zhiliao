package main

import "fmt"

func main() {
	// map定义 map[键类型]值类型
	// map 结构不能使用 cap()
	var m map[string] string
	fmt.Println(m)	// map[]
	fmt.Println(len(m))	// 0

	m1 := make(map[int] string, 3)	// 3 表示 map 的容量
	fmt.Println(m1)

	// 赋值
	m1[1] = "loedan0"
	m1[2] = "loedan1"
	m1[3] = "loedan2"
	m1[4] = "loedan3"
	fmt.Println(m1)	// map 中数据是无序的 map[1:loedan0 2:loedan1 3:loedan2 4:loedan3]

	// 定义并初始化
	m2 := map[string] string {"name": "loedan", "job": "golang"}
	fmt.Println(m2)		// map[job:golang name:loedan]
}
