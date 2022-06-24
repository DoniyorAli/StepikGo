package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Print("Enter the name: ")
	fmt.Scan(&name)

	fmt.Print("Enter the age: ")
	fmt.Scan(&age)

	fmt.Println(name, age)

}