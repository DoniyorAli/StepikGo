package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	_  // skip 7
	Add
)

func main() {

	fmt.Println(Sunday)   // вывод: 0
	fmt.Println(Saturday) // вывод: 6
	fmt.Println(Add) // вывод: 8

}