package main

import(
	"fmt"
	"code.oldboy.com/day1/assignments/content"
	"time"
)

func main() {
	var oldboy string
	var result string
	oldboy = "打印字符串"
	result = content.PrintStr(oldboy)
	fmt.Printf("%s", result)
	time.Sleep(10*time.Second)
}