package main

import(
	"fmt"
	"time"
)

//定义一个奇数的数组
func PrintOdd(odd int) {
	// var number int
	fmt.Printf("你选择的%d以内的奇数为: \n", odd)	
	for number := 0; number < odd; number++ {
		if number % 2 == 1 {
			fmt.Printf("%d \n", number)
		}
	}
}

//定义一个偶数的数组
func PrintEven(even int) {
	// var number int
	fmt.Printf("你选择的%d以内的偶数为: \n", even)		
	for number := 0; number < even; number++ {
		if number % 2 == 0 {
			fmt.Printf("%d \n", number)
		}
	}
}

func main() {
	go PrintOdd(10)
	go PrintEven(10)

	time.Sleep(10*time.Second)
}