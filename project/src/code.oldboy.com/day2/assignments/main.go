// package main

// import(
// 	"fmt"
// 	// "math"
// 	"math/rand"
// )
// //素数
// func isPrime(n int) bool {
// 	if n <=1 {
// 		return false
// 	}
// 	for i:=2; i < n; i++ {
// 		if n % i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// //水仙花
// // func isNarcissistic(n int) bool {
// // 	if n < 100 || n > 999{
// // 		return false
// // 	}
// // 	i,j,k = float64(n/100),float64n/10 %10),float64(n %10)
// // 	if math.Pow(i,3)+math.Pow(j,3)+math.Pow(k,3) == float64(n) {
// // 		return true
// // 	}
// // 	return false
// // }

// //阶乘
// func factorial(n int) int{
// if n <= 0 {
// 	return -1
// }
// var cnt int = 1
// for i:=2;i<=n;i++{
// 	cnt = cnt * i
// }
// 	return cnt
// }
// //递归
// func factorialRecursion(n int) int{
// 	if n <=0 {
// 		return -1
// 	}
// 	return 1
// }
// //生成伪随机数
// func random() {
// 	for i:=0;i<10;i++{
// 		fmt.Println(rand.Int())
// 	}
// }

// func main() {
// 	//调用素数函数
// 	for i:=101;i<=200;i++ {
// 		if isPrime(i){
// 			fmt.Println(i)
// 		}
// 	}
// 	//调用水仙花函数
// 	// for j:=100;j<999;j++ {
// 	// 	if isNarcissistic(j) {
// 	// 		fmt.Println(j)
// 	// 	}
// 	// }
// }