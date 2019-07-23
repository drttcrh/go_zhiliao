package _1_basic

import "fmt"

// 校验函数
func validate(name string, pwd string, email string) bool {
	return true
}

// 发送邮件函数
func emailSend(email string) {
	// todo 发送邮件
}

// 注册函数
func register(name string, pwd string, email string) string {
	validateRes := validate(name, pwd, email)
	if !validateRes {
		return "参数校验失败"
	}
	emailSend(email)
	return "注册成功"
}

func main() {
	// 网站注册流程练习
	result := register("loedan", "123456", "loedan@go.com")
	fmt.Println(result)
}
