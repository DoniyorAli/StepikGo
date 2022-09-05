package main     // BREAK with for loop

import "fmt"

func main() {

	var sum = 0

	for i := 1; i <= 9; i++ {
		if i > 4 {
			break
		}
		sum += i
	}
	fmt.Println("Total: ", sum)

}
