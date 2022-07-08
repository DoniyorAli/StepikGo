package main

import "fmt"

func main() {

	var num int32
	fmt.Scan(&num)
	num = num / 10 % 10
	fmt.Println(num)

}