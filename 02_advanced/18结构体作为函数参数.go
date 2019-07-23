package main

import "fmt"

type tea struct {
	name string
	color string
}

// 结构体作为函数参数
func structParams(t tea) {
	t.color = "蓝色"
	fmt.Println(t)		// {绿茶 蓝色}
}

func main() {
	t := tea{"绿茶", "绿色"}
	// 值传递
	structParams(t)
	fmt.Println(t)		// {绿茶 绿色}
}
