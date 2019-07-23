package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 创建文件
	fp, err := os.Create("./src/learn_go_zhiliao/05_use_pkg/b.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	// WriteString() 方法写数据入文件
	fp.WriteString("hello world\n")	// unix \n 为换行，win系统使用 \n\r
	fp.WriteString("性感荷官在线发牌\n")


	// Write() 方法写数据
	n, err1 := fp.Write([]byte{'l', 'o', 'e', 'd', 'a', 'n'})
	if err1 != nil {
		fmt.Println("write方法写入文件失败")
		return
	} else {
		fmt.Println(n)
	}


	// 使用光标流获取光标位置
	// io.SeekEnd，文件末尾位置
	count, _ := fp.Seek(0, io.SeekEnd)
	fmt.Println(count)
	fp.WriteAt([]byte{'l', 'o', 'e', 'd', 'a', 'n'}, count)


	defer fp.Close()	// 关闭文件，及时收回资源
}
