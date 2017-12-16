package main

import (
	"fmt"
	"strings"
)
// 定义一个函数，接受的参数名为fuffix,返回一个函数
func makeSuffixFunc(suffix string) func(string) string {
	alias_func := func(name string) string {
		//判断返回函数传入的参数是否为根函数传入的参数,如果不是，则添加stuffix的后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		//如果后缀存在，则直接返回
		return name
	}
	return alias_func
	//接收一个参数为name的参数
	// return func(name string) string {
	// 	//判断返回函数传入的参数是否为根函数传入的参数,如果不是，则添加stuffix的后缀
	// 	if !strings.HasSuffix(name, suffix) {
	// 		return name + suffix
	// 	}
	// 	//如果后缀存在，则直接返回
	// 	return name
	// }
}

func main() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}