package main

import (
	"fmt"
)

func modify_arr(a *[6]int) {
	fmt.Printf("modify:%p\n",a)
	a[0] = 100
}

func test6() {
	var b *[6]int
	modify_arr(b)
	fmt.Printf("After the modify,the value is: %d", b)
}


func main() {
	test6()
}