package main

import (
	"fmt"
)

func calc(n int) int {
	if n == 1 {
		return 1
	}
	return calc(n-1)*n
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	// n := calc(5)
	// fmt.Println(n)
	for i := 0;i < 10;i++ {
		n := fab(i)
		fmt.Println(n)
	}
}