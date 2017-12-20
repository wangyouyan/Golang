package main

        import "fmt"

        // func main() {
        //     var a []int = make([]int,5,10)
        //     a[4] = 100
        //     b := a[0:10]
        //     b[9] = 100
		// 	fmt.Printf("%v", a)
		// 	fmt.Println("\n")
        //     fmt.Printf("%v", b)
		// }
		
		// func main() {
		// 	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
		// 	b := ar[5:7]
		// 	fmt.Printf("%v\n",b)
		// 	fmt.Print(len(b))
		// 	fmt.Print(cap(b))
		// 	a := ar[0:4]
		// 	fmt.Printf("%v\n",a)
		// 	fmt.Print(len(a))	
		// 	fmt.Print(cap(a))			
		// }

		func main() {
			var a []int = make([]int, 5, 10)
            a[4] = 100
            b := a[0:10]
            b[9] = 100
            fmt.Printf("%v\n", a)
            fmt.Printf("%v\n", b)
		}