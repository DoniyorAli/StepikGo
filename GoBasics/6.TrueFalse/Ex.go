package main

import "fmt"

func main() {
//  ! --> negation
	var a bool = true
	var b bool = !a // false
	var c bool = !b // true
	fmt.Println(c)
//_____________________________________________________________________
//  && --> and   kamida bitta o'zgaruvchi notog'ri bo'lsa false aytaradi!
	box1 := 11
	box2 := 9
	bag := box1 > box2 && box1 < box2
	fmt.Println(bag) // --> false
//_____________________________________________________________________
//  || --> or    kamida bitta o'zgaruvchi tog'ri bo'lsa true aytaradi!
	box3 := 11
	box4 := 9
	box5 := 7
	box6 := 4
	box7 := true
	box8 := false

	bag1 := box3 > box4 || box6 < box5
	fmt.Println(bag1) // --> true

// Ikkala operand ham noto'g'ri bo'lsa, false qaytaradi!
	bag2 := box3 < box4 || box6 > box5
	fmt.Println(bag2) // --> false

	bag3 := box7 == box8 || box7 != box8
	fmt.Println(bag3) // --> true

}
