package main

import(
	"fmt"
	"os"
)

func swap(a int, b int) (int, int) {
	a = 3
	b = 4
	fmt.Printf("交换之前a的值为:%d, b的值为:%d", a, b)	
	a,b = b,a
	fmt.Printf("交换之前a的值为:%d, b的值为:%d", a, b)
	return a, b
}

// 交换两个数字的简写
func swapv2(a int, b int) (int, int) {
	return b, a 
}

func main() {
	name,err := os.Hostname()
	fmt.Printf("当前主机名为: %s \n错误日志输出: %v \n", name, err)

	val := os.Getenv("PATH")
	fmt.Printf("%s \n", val)

	swap(3,4)
}