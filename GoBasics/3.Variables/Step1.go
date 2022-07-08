package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)
	box1 := num * 2
	box2 := box1 + 100
	fmt.Println("Result:", box2)

}