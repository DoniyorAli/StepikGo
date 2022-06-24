package main

import "fmt"

func main() {

	// gradus := 360
	// hour := 12
	// clock := gradus / hour  // --> 30 gradus --> 30 = 1 hour

	var num int
	fmt.Scan(&num)
	hour := num / 30
	minute := num % 30 * 2
	fmt.Println("It is",hour,"hours",minute,"minutes")

}