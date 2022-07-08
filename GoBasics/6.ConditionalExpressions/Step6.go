package main

import "fmt"

func main() {

	year := 0
	fmt.Scanln(&year)
 
	if year % 400 == 0 {
		fmt.Println("YES")
	}else if year % 4 == 0 && year % 100 != 0{
		fmt.Println("YES")
	}else if year % 100 == 0 || year % 100 != 0{
		fmt.Println("NO")
	}

	
	// leap := (year % 4 == 0) && (year % 100 != 0) || (year % 400 == 0)
	// if leap {
	// 	fmt.Println("YES")
	// }else {
	// 	fmt.Println("NO")
	// }

}