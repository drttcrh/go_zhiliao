package main

import "fmt"

func main() {

	// if 嵌套的使用
	score := 600
	if score >= 680 {
		fmt.Println("清华")
		if score > 800 {
			fmt.Println("蓝翔挖掘机")
		} else if score > 700 {
			fmt.Println("新东方厨师")
		} else {
			fmt.Println("苦逼程序员")
		}
	} else {
		fmt.Println("北大")
	}
}
