package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Open() 只读方式打开文件
	fp, err := os.Open("./src/learn_go_zhiliao/05_use_pkg/a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	fp.WriteString("this is lesson file")

	defer fp.Close()


	// os.OpenFile(文件名，打开方式，打开权限) 多模式打开文件
	fp1, err1 := os.OpenFile("./src/learn_go_zhiliao/05_use_pkg/b.txt", os.O_RDWR, 6)
	if err1 != nil {
		fmt.Println("open file failed!")
	}
	fp1.WriteString("ahaschool")
	defer fp1.Close()
}
