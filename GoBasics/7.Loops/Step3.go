package main

import "fmt"

func main() {

	var n int 
	fmt.Scan(&n)
	m := 0
	total := 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		if (9 < m && m < 99) && (m % 8 == 0) {
			total += m
		}
	}
	fmt.Println(total)
}