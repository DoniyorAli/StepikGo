package main

import "fmt"

func main() {

	var number uint16
	fmt.Scan(&number)

	if 1 <= number && number < 10{
		fmt.Println(number)
	}else if 10 <= number && number < 100{
		number := number / 10
		fmt.Println(number)
	}else if 100 <= number && number < 1000{
		number := number / 100
		fmt.Println(number)
	}else if 1000 <= number && number < 10000{
		number := number / 1000
		fmt.Println(number)
	}else if number == 10000{
		number := number / 10000
		fmt.Println(number)
	}else {
		fmt.Println("Enter from 1 to 10.000!")
	}

	// switch {
	// case 1 <= number && number < 10: fmt.Println(number)
	// case 10 <= number && number < 100:
	// 	fmt.Println(number / 10)
	// case 100 <= number && number < 1000:
	// 	fmt.Println(number / 100)
	// case 1000 <= number && number < 10000:
	// 	fmt.Println(number / 1000)
	// case number == 10000:
	// 	fmt.Println(number / 10000)
	// }

	
	// package main

	// import "fmt"

	// func main() {
	// 	var a int
	// 	fmt.Scan(&a)
	// 	fmt.Println(getFirstNumber(a))
	// }

	// func getFirstNumber(a int) int {
	// 	if a < 10 {
	// 		return a
	// 	} else {
	//  		return getFirstNumber(a / 10)
	// 	}	
	// }
}