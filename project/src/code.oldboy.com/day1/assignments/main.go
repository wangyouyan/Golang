package main

import(
	"fmt"
	"code.oldboy.com/day1/assignments/content"
	"time"
)

func main() {
	var str string
	var binary int
	var floatf float64
	var decim int
	var hex int

	str = content.PrintStr("打印字符串")
	binary = content.PrintBinary(10)
	floatf = content.PrintFloating(78.9)
	decim = content.PrintDecimalism(10)
	hex = content.PrintSHexadecimal(465)

	fmt.Println("本周作业,各种格式输出\n")
	time.Sleep(1*time.Second)
	fmt.Printf("2017-12-01 %a \n", str)
	time.Sleep(1*time.Second)	
	fmt.Printf("二进制输出: %b \n", binary)
	time.Sleep(1*time.Second)	
	fmt.Printf("浮点数 %f \n", floatf)
	time.Sleep(1*time.Second)	
	fmt.Printf("十进制%d \n", decim)
	time.Sleep(1*time.Second)	
	fmt.Printf("十六进制 %x \n", hex)
}