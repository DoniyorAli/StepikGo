package main     // CONTINUE with loop

import "fmt"

func main() {

	var sum = 0
 
	for i := 1; i <= 10; i++{
    if i % 2 == 0 {
        continue        // переходим к следующей итерации
    }
    	sum += i
	}
	fmt.Println("[from 1 to 10] an odd numbers total:",sum)
}