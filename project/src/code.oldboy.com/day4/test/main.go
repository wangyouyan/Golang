package main

import (
	"fmt"
)

func add(a,b int)(c int) {
	c = a + b
	return 
}

func calc(a,b int)(sum int,avg int) {
	sum = a + b
	avg = (a+b)/2
	return
}

func main() {
	fmt.Printf("%d\n",add(1,2))
	fmt.Printf("%d,%d\n",calc(5,6))
}