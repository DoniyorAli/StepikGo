package main

import "fmt"

const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	e
	f = iota + 5
	_
	t
	l = e - 2
	_
	i
	j
	k = iota + 2
	g
	r = t + 7
	o
	p
	x = iota + 6
)

func main() {

	fmt.Println(a) // --> 1
	fmt.Println(d) // --> 6
	fmt.Println(f) // --> 11
	fmt.Println(l) // --> 4
	fmt.Println(k) // --> 15
	fmt.Println(r) // --> 20
	fmt.Println(x) // --> 24

}