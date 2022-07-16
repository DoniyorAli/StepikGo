package main

import "fmt"

func main() {

	A := 0
	B := 0
	total := 0
	for fmt.Scan(&A, &B); A <= B; A += 1 {
		total += A
	}
	fmt.Println(total)
//_______________________________________________________

	// I := 0
	// J := 0
	// Total := 0
	// fmt.Print("Enter first number (I<J): I = ")
	// fmt.Scan(&I)
	// fmt.Print("Enter second number: J = ")
	// fmt.Scan(&J)

	// for ; I <= J; I += 1{
	// 	Total += I
	// }
	// fmt.Println(Total)
//______________________________________________________

// Simple example with arithmetic

	// num1 := 0
	// num2 := 0
	// fmt.Print("Eter two numbers: ")
	// fmt.Scan(&num1, &num2)
	// totall := (num2 - num1 +1) * ((num1 + num2) / 2)
	// fmt.Println("Total: ", totall)
}