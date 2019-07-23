package main

import (
	"fmt"
	"strings"
)

// 查找一个字符串在另一个字符串中是否出现
func useContains()  {
	str := "hello"
	str1 := "h"
	b := strings.Contains(str, str1)
	if b {
		fmt.Printf("%s中包含%s\n", str, str1)
	} else {
		fmt.Printf("%s中不包含%s\n", str, str1)
	}
}

// 字符串连接
func useJoin()  {
	slice := []string{"loedan", "ahaschool", "golang"}
	str := strings.Join(slice, ",")
	fmt.Println(str)
}

// 查找一个字符串在另一个字符串中第一次出现的位置，返回第一次出现的下标， -1 表示未找到
func useIndex() {
	str := "hello"
	str1 := "h"
	i := strings.Index(str, str1)
	fmt.Println(i)
}

// 将一个字符串重复 n 次
func useRepeat() {
	str := "性感网友，在线取名。"
	str1 := strings.Repeat(str, 5)
	fmt.Println(str1)
}

// 字符串替换
func useReplace() {
	str := "性感网友在线取名性感"
	str1 := strings.Replace(str, "性感", "苗条", 1)	// 参数 4 为 -1 时， 将所有匹配到的全都替换
	fmt.Println(str1)
}

// 切割字符串
func useSplit() {
	str := "133-2222-1111"
	str1 := strings.Split(str, "-")
	fmt.Println(str1)
	str2 := strings.Join(str1, "")
	fmt.Println(str2)
}

// 删除字符串头尾内容
func useTrim() {
	str := "===ab==c&*$=="
	str1 := strings.Trim(str, "=")
	fmt.Println(str1)
}

// 去除字符串中的空格，转成切片，一般用于统计单词的个数
func useFields() {
	str := "   are you ok    "
	slice := strings.Fields(str)
	fmt.Println(slice)
}

func main() {

	// 查找一个字符串在另一个字符串中是否出现
	useContains()

	// 字符串连接
	useJoin()

	// 查找一个字符串在另一个字符串中第一次出现的位置，返回第一次出现的下标， -1 表示未找到
	useIndex()

	// 将一个字符串重复 n 次
	useRepeat()

	// 字符串替换
	useReplace()

	// 切割字符串
	useSplit()

	// 删除字符串头尾内容
	useTrim()

	// 去除字符串中的空格，转成切片，一般用于统计单词的个数
	useFields()
}
