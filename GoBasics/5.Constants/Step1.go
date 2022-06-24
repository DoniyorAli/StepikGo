package main

import "fmt"

const pi float64 = 3.1415  // The const keyword is used to define constants
//We cannot change the value of a constant
// pi = 2.7182 // ! Error

const (   // Constants, like ordinary variables, can be declared in a block:
	a int = 45
	b float32 = 3.3
)

/*You can also leave out the value of the 
following constant in order (the value will be copied):
*/
const( 
	A int = 45
	B
	C float32 = 3.3
	D
)

func main() {

	fmt.Println(A, B, C, D)  // Вывод: 45 45 3.3 3.3

}