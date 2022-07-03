package main

import "fmt"

func main() {

	number := 0
	fmt.Scan(&number)

	first := number / 10 / 10
	second := number / 10 % 10
	third := number % 10

	if first == second || second == third{
		fmt.Println("NO")
	}else if first != second || second != third{
		fmt.Println("YES")
	}
}