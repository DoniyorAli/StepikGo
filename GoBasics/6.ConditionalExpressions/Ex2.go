package main

import "fmt"

func main() {

	var a int = 11
    var b int = 7
	
    var c bool = a == b
    fmt.Println(c)      // false

	var d bool = a > b   // true
	fmt.Println(d)

	var f bool = a < b   // false
	fmt.Println(f)

	var e bool = a <= b  // false
	fmt.Println(e)

	var j bool = a >= b  // true
	fmt.Println(j)
	
	fmt.Println()
	var k bool = a != b // true
	var l bool = a != 11 // false
	fmt.Println(k, l)

}