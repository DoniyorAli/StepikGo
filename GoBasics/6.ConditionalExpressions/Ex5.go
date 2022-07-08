package main

import "fmt"

func main() {

	var c uint32
	fmt.Println("Enter numbers until from 1... to 300!:")
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 50:
		fmt.Println("от 1 до 50")
	case 51 <= c && c <= 100:
		fmt.Println("from 51 to 100")
	case 101 <= c && c <= 150:
		fmt.Println("from 101to 150")
	case 151 <= c && c <= 200:
		fmt.Println("from 151to 200")
	case 201 <= c && c <= 250:
		fmt.Println("from 201to 250")
	case 251 <= c && c <= 300:
		fmt.Println("from 251to 300")
	default: fmt.Println("Wrong input!")
	}
}