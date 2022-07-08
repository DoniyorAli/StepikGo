package main

import "fmt"

func main() {

	var number int32
	fmt.Scanln(&number)

	box1 := number / 1000
	total1 := (box1 / 100) + (box1 % 10) + (box1 / 10 % 10)
	fmt.Println(total1)

	box2 := number % 1000
	total2 := (box2 / 100) + (box2 / 10) % 10 + (box2 % 10)
	fmt.Println(total2)

	if total1 == total2{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
	
}