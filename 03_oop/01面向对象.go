package main

import "fmt"

// 人结构体
// 结构体嵌套结构体
// 父类
type Person1 struct {
	id int
	name string
	age int
}

// 学生结构体
// 子类
type Student1 struct {
	Person1	// 匿名字段
	score int
}

// 创建对象
func createObject() {

	// 顺序初始化
	var s Student1 = Student1{Person1{1, "loedan", 18}, 100}
	fmt.Println(s)		// {{1 loedan 18} 100}

	var p Person1 = Person1{1, "harden", 12}
	fmt.Println(p)		// {1 harden 12}

	// 自动推导方式
	s2 := Student1{Person1{2, "jordan", 13}, 90}
	fmt.Println(s2)		// {{2 jordan 13} 90}

	// 指定成员初始化，未填初始化值的元素会赋予该类型的默认值
	s3 := Student1{score: 98}
	fmt.Println(s3)		// {{0  0} 98}

	s4 := Student1{Person1: Person1{3, "aha", 23}}
	fmt.Println(s4)		// {{3 aha 23} 0}
}

// 成员操作
func objOpt() {

	s := Student1{Person1{1, "loedan", 13}, 90}
	fmt.Println(s)
	s.age = 18
	fmt.Println(s)
	s.Person1.age = 28
	fmt.Println(s)
	s.Person1 = Person1{2, "aha", 23}
	fmt.Println(s)
	s.score = 101
	fmt.Println(s)
}

func main() {
	// 创建对象
	createObject()

	// 成员操作
	objOpt()
}
