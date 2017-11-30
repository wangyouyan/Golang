package main

import(
	"fmt"
)

func main() {
	//chan 表示管道
	//string 表示管道类型
	//数字3 表示管道的最大长度为3
	pipe := make(chan string, 3)

	// 把类型为strint的"a"传给管道
	pipe <- "a"
	pipe <- "b"
	// 定义一个为"a"的字符串,然后把管道中的变量赋值给"a"
	var a string
	a = <- pipe

	fmt.Printf("当前管道值为: %s\n", a)
}