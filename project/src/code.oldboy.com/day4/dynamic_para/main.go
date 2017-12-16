package main


import (
	"fmt"
)

func Add(arg...int) {
	var sum int
	for i := 0;i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return sum 
}

func testArg() {
	fmt.Println(Add(1))
	fmt.Println(Add())
	fmt.Println(Add(1,2,3,4,5))
}

func main() {
	testArg()
}