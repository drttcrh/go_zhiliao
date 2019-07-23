package main

import (
	"fmt"
	"os"
)

func main() {

	// 创建文件
	// 参数需要路径，支持绝对路径和相对路径
	// 当前路径默认为 GOPATH 指定的路径
	fp, err := os.Create("./src/learn_go_zhiliao/05_use_pkg/a.txt")
	if err != nil {
		// 创建失败可能原因
		// 1、路径不存在
		// 2、文件/目录权限不足
		// 3、程序打开文件上限
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()	// 关闭文件，及时收回资源
}
