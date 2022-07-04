package main

import "fmt"

func main() {

	number := 123456

	box1 := number % 1000
	// box2 := number / 1000
	
	box1 = box1 % 10
	box1 = box1 % 10

	fmt.Println(box1)

}