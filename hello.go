package main // 程序的包名

import (  // 导入模块,多行导入使用括号
	"fmt"
	"time"
	)

// main函数
func main() {  // 函数的左花括号，必须要和函数名在同一行
	// golang中的表达式，末尾加不加";"都可以，建议不加
	fmt.Println("hello world")  // 打印字符串

	time.Sleep(1*time.Second)
}
