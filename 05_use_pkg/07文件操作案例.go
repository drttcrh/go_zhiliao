package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var sourceFile string
	var targetFile string
	fmt.Println("请输入读取的源文件名：")
	fmt.Scan(&sourceFile)
	fmt.Println("请输入新生成的文件名：")
	fmt.Scan(&targetFile)

	if sourceFile == targetFile {
		fmt.Println("读取的文件名和生成的文件名不能一样")
		return
	}

	// 读取文件内容
	fpr, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("源文件打开失败，请检查文件名或路径是否正确")
		return
	}
	defer fpr.Close()

	// 写入新文件
	fpw, err1 := os.Create(targetFile)
	if err1 != nil {
		fmt.Println("文件创建失败", err1)
		return
	}
	defer fpw.Close()

	buf := make([]byte, 1024 * 4)
	for {
		// 读取内容
		n, err := fpr.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
			} else {
				fmt.Println("文件读取失败")
			}
			return
		}
		// 写入内容
		fpw.Write(buf[:n])
	}
}
