package main

import "fmt"

func main() {

	num := 0
	max := 0
	count := 1

	for fmt.Scanln(&num); num != 0; fmt.Scanln(&num) {
		if num > max {
			max = num
			count = 1
		}else if num == max {
			count++
		}
	}
	fmt.Println(count)
}