package main

import "fmt"

func main() {

	var num1 int
	var num2 int

	fmt.Scan(&num1, &num2)

	num1 = num1 * num1
	num2 = num2 * num2
	c := num1 + num2

	fmt.Println(c)

}

// package main
// import "fmt"
// func main(){

// 	var a, b int
// 	fmt.Scan(&a, &b)

// 	fmt.Println((a*a) + (b*b))

// }