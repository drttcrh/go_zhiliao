package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func useOs() {
	// os.Open() 只读方式打开文件
	fp, err := os.Open("./src/learn_go_zhiliao/05_use_pkg/b.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	defer fp.Close()

	buf := make([]byte, 10)
	// 块读取
	n, err1 := fp.Read(buf)	// n 为读取的长度
	if err1 != nil {
		fmt.Println("read file failed")
	}
	fmt.Println(n)
	fmt.Println(string(buf[:n]))


	// 循环读取文件所有内容
	for {
		n1, err := fp.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n1]))
	}
}

func useBufio()  {
	fp, err := os.Open("./src/learn_go_zhiliao/05_use_pkg/b.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	defer fp.Close()

	// 创建一个缓冲区
	r := bufio.NewReader(fp)
	// 逐行读取
	slice, _ := r.ReadBytes('\n')
	fmt.Println(string(slice))
	slice, _ = r.ReadBytes('\n')
	fmt.Println(string(slice))


	// 循环读取所有内容
	for {
		buf, err1 := r.ReadBytes('\n')
		fmt.Println(string(buf))
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			fmt.Println(err1)
		}
	}
}

func main() {
	// 使用 os 包读取文件
	useOs()

	//
	useBufio()
}
